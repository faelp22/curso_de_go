package main

import (
	"fmt"
	"log"

	"github.com/faelp22/tcs_curso/stoq/config"
	"github.com/faelp22/tcs_curso/stoq/entity"
	"github.com/faelp22/tcs_curso/stoq/pkg/database"
)

func main() {

	config := config.NewConfig("8080", config.DBConfig{
		DB_DRIVE: "sqlite3",
		// DB_HOST:  "192.168.0.100",
		// DB_PORT:  "5432",
		// DB_USER:  "root",
		// DB_PASS:  "123456",
		DB_NAME: "db.sqlite3",
	}, false)

	fmt.Println(config.DB_DSN)

	dbpool := database.NewDB(config)

	stmt, err := dbpool.DB.Prepare("select id, name from tb_produto where id = ?")
	if err != nil {
		log.Println(err.Error())
	}

	defer stmt.Close()

	prod := entity.Produto{}

	err = stmt.QueryRow(5).Scan(prod.ID, prod.Name)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Printf("Produto ID %d Name %v \n", prod.ID, prod.Name)

	// p1 := entity.Produto{
	// 	ID:        1,
	// 	Name:      "PS5",
	// 	Code:      "ASDF",
	// 	Price:     5000,
	// 	CreatedAt: "16/05/22 19:00:00",
	// 	UpdatedAt: "16/05/22 19:00:00",
	// }

	// fmt.Println(p1.String())

}
