package main

import (
	"fmt"

	"github.com/faelp22/tcs_curso/stoq/config"
)

func main() {

	config := config.NewConfig("8080", config.DBConfig{
		DB_DRIVE: "sqlite",
		// DB_HOST:  "192.168.0.100",
		// DB_PORT:  "5432",
		// DB_USER:  "root",
		// DB_PASS:  "123456",
		DB_NAME: "stoq.sqlite3",
	}, false)

	fmt.Println(config.DB_DSN)

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
