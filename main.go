package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type headers struct {
	XForwardedProto string `json:"x-forwarded-proto"`
	XForwardedPort  string `json:"x-forwarded-port"`
	Host            string `json:"host"`
}

type JsonResponse struct {
	Headers headers
}

func main() {
	//Encode data:
	postBody, _ := json.Marshal(map[string]string{
		"name":  "Toby",
		"email": "Toby@example.com",
	})
	responseBody := bytes.NewBuffer(postBody)
	// http.Post Implementation:
	resp, err := http.Post("https://postman-echo.com/post", "application/json", responseBody)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}

func get() {
	resp, err := http.Get("https://postman-echo.com/get?foo1=bar1&foo2=bar2")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var jsonResponse JsonResponse

	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		panic(err)
	}

	fmt.Println(jsonResponse)
}

func newRequest() {
	// Statement yang menghasilkan instance http.Client, diperlukan untuk eksekusi request
	var client = &http.Client{}

	// http.NewRequest() digunakan untuk membuat request baru
	req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/ability/?limit=1", nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("status:", resp.Status)

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))

}
