package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/diegolopezcode/api-crud-complete-chi/configs"
)

// GetServer
// The function GetServer is a handler function that takes in a ResponseWriter and a Request object and
// returns a 200 status code with a message "Server is running"
func GetServer(w http.ResponseWriter, r *http.Request) bool {
	url := configs.Config("URL_BD_ZINSEARCH") + "ui/"
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	fmt.Println(url)

	if err != nil {
		fmt.Println(err)
		return false
	}
	req.Header.Add("Authorization", configs.Config("AUTHORIZATION"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(string(body))
	return true
}
