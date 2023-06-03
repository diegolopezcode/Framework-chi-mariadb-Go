package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/diegolopezcode/api-crud-complete-chi/configs"
)

// It reads a file, sends it to a server, and then prints the response

func SendFile(filename string, wg1 *sync.WaitGroup, contador1 *int64, w http.ResponseWriter) {
	fmt.Println("Enviando archivo: " + filename)
	atomic.AddInt64(contador1, 1)
	runtime.Gosched()
	fmt.Println("Contador", atomic.LoadInt64(contador1))

	url := configs.Config("URL_BD_ZINSEARCH") + "api/_bulk"
	method := "POST"
	dat, err := os.ReadFile(configs.Config("ROUTE_TEMP_FILES") + filename)
	if err != nil {
		Error_server(w, err)
		wg1.Done()
		return
	}

	payload := strings.NewReader(string(dat))
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		Error_server(w)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic YWRtaW46Q29tcGxleHBhc3MjMTIz")

	res, err := client.Do(req)
	if err != nil {
		Error_server(w)
		wg1.Done()
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		Error_server(w)
		return
	}
	fmt.Println(string(body))

	wg1.Done()
}

// It takes a ResponseWriter and sets the header to application/json, sets the status code to 404, and
// then encodes a map of strings to strings with the message "Servidor no encontrado"
func Error_server(w http.ResponseWriter, err ...error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Servidor no encontrado",
	})
}
