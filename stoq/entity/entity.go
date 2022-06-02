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
	CreatedAt string  `json:"created_at,omitempty"`
	UpdatedAt string  `json:"updated_at,omitempty"`
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

type ProdutoList struct {
	List []*Produto `json:"list"`
}

func (pl ProdutoList) String() string {
	data, err := json.Marshal(pl)

	if err != nil {
		log.Println("error to convert ProdutoList to JSON")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

func NewProduto(nome, code string, price float64) *Produto {
	return &Produto{
		Name:  nome,
		Code:  code,
		Price: price,
	}
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewAdmin() *User {
	return &User{
		Username: "admin",
		Password: "supersenha",
	}
}

type Token struct {
	Token string `json:"token"`
}

const USER_TOKEN = "fake-WzD5fqrlaAXLv26bpI0hxvAhDp7T1Bac"
