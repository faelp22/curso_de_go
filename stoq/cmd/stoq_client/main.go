package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/faelp22/browser/bro"
	"github.com/faelp22/tcs_curso/stoq/entity"
)

type BrowserCli interface {
	Get(url string) (*http.Response, error)
	Post(url string, payload io.Reader) (*http.Response, error)
	Put(url string, payload io.Reader) (*http.Response, error)
	Delete(url string) (*http.Response, error)
}

func main() {

	base_url := "http://localhost:8080"
	ssl_verify := false
	header := http.Header{
		"Content-Type": []string{"application/json; charset=utf-8"},
	}

	bro := bro.NewBrowser(base_url, ssl_verify, header, 10)

	token := Login(bro)

	bro.AddHeader("Authorization", "Bearer "+token.Token)

	resp, err := bro.Get("/api/v1/products")
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

func Login(bro BrowserCli) *entity.Token {
	url := "/api/v1/user/login"

	user := entity.NewAdmin()

	dados, _ := json.Marshal(user)

	// fmt.Println(string(dados))

	payload := bytes.NewBuffer(dados)

	resp, err := bro.Post(url, payload)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler Body da request")
		fmt.Println(err.Error())
	}

	fmt.Println(string(body))

	token := entity.Token{}

	err = json.Unmarshal(body, &token)
	if err != nil {
		fmt.Println("Erro ao ler Body da request")
		fmt.Println(err.Error())
	}

	// fmt.Println(token)

	return &token
}
