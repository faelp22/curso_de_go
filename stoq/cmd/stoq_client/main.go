package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/faelp22/browser"
	"github.com/faelp22/browser/prepare"
	"github.com/faelp22/tcs_curso/stoq/entity"
	"github.com/faelp22/tcs_curso/stoq/pkg/crypto"
)

func main() {

	bro := browser.NewBrowser(browser.BrowserConfig{
		BaseURL:   "http://localhost:8080",
		SSLVerify: false,
		Header: http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		},
		Timeout: 3, // 3 segundos
	})

	token := Login(bro)

	bro.SetHeader("Authorization", "Bearer "+token.Token)

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

	ok := crypto.WriteCryptoFile("./dados.txt", body)
	if ok {
		fmt.Println("Arquivo criado")
	}

	fmt.Println(string(body))
}

func Login(bro browser.BrowserCli) *entity.Token {
	url := "/api/v1/user/login"

	user := entity.NewAdmin()

	dados, _ := json.Marshal(user)

	// fmt.Println(string(dados))

	payload := prepare.PrepareJSON(bro, dados)

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

	// fmt.Println(string(body))

	token := entity.Token{}

	err = json.Unmarshal(body, &token)
	if err != nil {
		fmt.Println("Erro ao ler Body da request")
		fmt.Println(err.Error())
	}

	// fmt.Println(token.Token)

	return &token
}
