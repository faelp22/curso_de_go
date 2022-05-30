package main

import (
	"encoding/json"
	"fmt"

	"github.com/faelp22/tcs_curso/stoq/config"
	"github.com/faelp22/tcs_curso/stoq/pkg/database"
	"github.com/faelp22/tcs_curso/stoq/pkg/service"
)

func main() {

	config := config.NewConfig("8080", config.DBConfig{
		DB_DRIVE: "sqlite3",
		// DB_HOST:  "192.168.0.100",
		// DB_PORT:  "5432",
		// DB_USER:  "root",
		// DB_PASS:  "123456",
		DB_NAME: "db.sqlite3",
	})

	dbpool := database.NewDB(config)
	service := service.NewProdutoService(dbpool)

	lista_produtos := service.GetAll()

	dados, err := json.Marshal(lista_produtos)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(dados))
}
