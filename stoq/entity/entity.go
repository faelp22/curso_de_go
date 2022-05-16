package entity

import (
	"encoding/json"
	"log"
)

type ProdutoInterface interface {
	String() string
}

type Produto struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Code      string  `json:"code"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

func (p Produto) String() string {
	data, err := json.Marshal(p)

	if err != nil {
		log.Println("error to convert Produto to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}