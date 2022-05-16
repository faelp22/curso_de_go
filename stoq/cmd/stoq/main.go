package main

import (
	"fmt"

	"github.com/faelp22/tcs_curso/stoq/entity"
)

type Computador string

func (c Computador) Ligar() string {
	return "Ligando ....."
}

func main() {

	p1 := entity.Produto{
		ID:        1,
		Name:      "PS5",
		Code:      "ASDF",
		Price:     5000,
		CreatedAt: "16/05/22 19:00:00",
		UpdatedAt: "16/05/22 19:00:00",
	}

	var pc1 Computador

	pc1 = "Dell"

	fmt.Println(pc1.Ligar())

	fmt.Println(p1.String())

}
