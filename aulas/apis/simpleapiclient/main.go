package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	response, err := http.Get("http://localhost:3000/user")
	if err != nil {
		fmt.Println(err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))

	response, err = http.PostForm("http://localhost:3000/user", url.Values{"Nome": {"Davi"}})
	if err != nil {
		fmt.Println(err)
	}
	data, _ = ioutil.ReadAll(response.Body)
	fmt.Println(string(data))

	type User struct {
		Name string `json:"name"`
	}
	user := "Davi"
	userJson, _ := json.Marshal(user)
	response, err = http.Post("http://localhost:3000/user", "application/json", bytes.NewBuffer(userJson))
	if err != nil {
		fmt.Println(err)
	}
	data, _ = ioutil.ReadAll(response.Body)
	fmt.Println(string(data))

	req, err := http.NewRequest("POST", "http://localhost:3000/admin", bytes.NewBuffer(userJson))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	response, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	data, _ = ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
}
