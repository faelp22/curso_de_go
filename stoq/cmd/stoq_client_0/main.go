package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	url := "http://localhost:8080/api/v1/products"

	http_client := &http.Client{}

	http_req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Erro ao criar HTTP GET Request")
		fmt.Println(err.Error())
	}

	resp, err := http_client.Do(http_req)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler Body da request")
		fmt.Println(err.Error())
	}

	if resp.StatusCode > 200 {
		fmt.Println("Erro na requisição")
		fmt.Printf("StatusCode: %v\n", resp.StatusCode)
		fmt.Println(string(body))
		os.Exit(1)
	}

	fmt.Println(string(body))

}
