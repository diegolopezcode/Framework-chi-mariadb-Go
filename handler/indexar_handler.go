package handler

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/diegolopezcode/api-crud-complete-chi/configs"
	"github.com/diegolopezcode/api-crud-complete-chi/handler/common"
	"github.com/go-chi/chi/v5"
)

type Emails struct {
	From       string `json:"From"`
	To         string `json:"To"`
	Cc         string `json:"Cc"`
	Subject    string `json:"Subject"`
	Body       string `json:"Body"`
	Category   string `json:"Category"`
	Date       string `json:"Date"`
	X_from     string `json:"X_from"`
	X_to       string `json:"X_to"`
	X_cc       string `json:"X_cc"`
	X_bcc      string `json:"X_bcc"`
	X_folder   string `json:"X_folder"`
	X_origin   string `json:"X_origin"`
	X_filename string `json:"X_filename"`
}

type detIndex struct {
	DtIndex string `json:"_index"`
}
type genIndex struct {
	Index detIndex `json:"index"`
}

func Indexar(w http.ResponseWriter, r *http.Request) {

	con := GetServer(w, r)
	fmt.Println(con)

	paths := chi.URLParam(r, "paths")
	fmt.Println(paths)
	url := configs.Config("NFS_PATH") + paths + "/maildir"
	if _, err := os.Stat(url); os.IsNotExist(err) {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "No se encontro el directorio",
		})
		return
	}
	cd := 0
	recursiveSearch(url, url, &cd, paths, w)
	archivos, err := ioutil.ReadDir("temp")
	if err != nil {
		fmt.Println("Error al leer directorio")
	}

	var contador1 int64
	const gs1 = 1
	var wg1 sync.WaitGroup
	wg1.Add(len(archivos))
	for _, archivo := range archivos {

		go common.SendFile(archivo.Name(), &wg1, &contador1, w)
	}
	wg1.Wait()
	errr := os.RemoveAll("temp/")
	if errr != nil {
		fmt.Println("Error al eliminar directorio")
	}
	if con {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Indexacion exitosa",
		})
	}
}

func recursiveSearch(path, Category string, cd *int, paths string, w http.ResponseWriter) {

	archivos, err := ioutil.ReadDir(path)
	if err != nil {

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusRequestTimeout)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Error al leer el archivo",
		})
		return
	}
	for _, archivo := range archivos {

		if archivo.IsDir() {

			recursiveSearch(path+"/"+archivo.Name(), archivo.Name(), cd, paths, w)

		} else {

			if archivo.Size() > 0 {
				var contador int64
				var wg sync.WaitGroup
				wg.Add(1)

				*cd++
				go readFile(path+"/"+archivo.Name(), Category, &wg, &contador, cd, paths)
				wg.Wait()
			}

		}

	}

}

func readFile(path, Category string, wg *sync.WaitGroup, contador *int64, cd *int, paths string) {

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Archivo ilegible")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	email := Emails{}
	flagBody := false
	flagSubject := false
	flagFrom := false
	for scanner.Scan() {
		var flagSubject2 *bool = &flagSubject
		var flagFrom2 *bool = &flagFrom
		a := fmt.Sprintln(scanner.Text())

		switch expression := a; {
		case strings.Contains(expression, "From:") && !strings.Contains(expression, "X-From:") && !flagSubject:
			email.From = strings.TrimSpace(strings.TrimPrefix(expression, "From: "))
			*flagFrom2 = true
		case strings.Contains(expression, "Cc:") && !strings.Contains(a, "X-Cc:") && !flagSubject:
			email.Cc = strings.TrimSpace(strings.TrimPrefix(expression, "Cc: "))
		case strings.Contains(expression, "Subject:") && !flagSubject:
			email.Subject = strings.TrimPrefix(expression, "Subject: ")
			*flagSubject2 = true
		case strings.Contains(expression, "Date:"):
			email.Date = strings.TrimPrefix(expression, "Date: ")
		case strings.Contains(expression, "X-From:"):
			email.X_from = strings.TrimPrefix(expression, "X-From: ")
		case strings.Contains(expression, "X-To:"):
			email.X_to = strings.TrimPrefix(expression, "X-To: ")
		case strings.Contains(expression, "X-Cc:"):
			email.X_cc = strings.TrimPrefix(expression, "X-Cc: ")
		case strings.Contains(expression, "X-bCc:"):
			email.X_bcc = strings.TrimPrefix(expression, "X-bCc: ")
		case strings.Contains(expression, "X-Folder:"):
			email.X_folder = strings.TrimPrefix(expression, "X-Folder: ")
		case strings.Contains(expression, "X-Origin:"):
			email.X_origin = strings.TrimPrefix(expression, "X-Origin: ")
		case strings.Contains(expression, "X-FileName:"):
			email.X_filename = strings.TrimPrefix(expression, "X-FileName: ")
			flagBody = true
		}

		if !flagSubject && flagFrom {
			if !strings.Contains(a, "From:") {
				email.To = strings.TrimSpace(email.To + strings.TrimSpace(strings.TrimPrefix(a, "To: ")))
			}
		}

		if flagBody {
			if !strings.Contains(a, "X-FileName:") {
				email.Body = email.Body + a

			}
		}
		email.Category = Category
	}

	bs, err := json.Marshal(email)
	if err != nil {
		fmt.Println("Error", err)
	}

	count := math.RoundToEven(float64(*cd / 3000))
	valorCount := fmt.Sprintf("%v", count)
	Indexacion(string("temp/"+valorCount+".ndjson"), bs, paths, wg, contador)

}

func Indexacion(filename string, data []byte, nameProject string, wg *sync.WaitGroup, contador *int64) {
	atomic.AddInt64(contador, 1)
	runtime.Gosched()
	var indexar genIndex
	indexar.Index.DtIndex = nameProject
	bs, err := json.Marshal(indexar)
	if err != nil {
		fmt.Println("Error", err)
	}

	if !fileExists(filename) {
		if err := os.Mkdir("temp", os.ModePerm); err != nil {
			fmt.Println(filename)
		}
		_ = ioutil.WriteFile(filename, bs, 0644)

	} else {
		AppendFile(filename, []byte("\n"))
		AppendFile(filename, bs)
	}
	AppendFile(filename, []byte("\n"))
	AppendFile(filename, data)
	wg.Done()

}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func AppendFile(filename string, data []byte) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.Write(data); err != nil {
		return err
	}

	return nil
}
