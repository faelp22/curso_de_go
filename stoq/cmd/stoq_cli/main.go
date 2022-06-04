package main

import (
	"fmt"

	"github.com/faelp22/tcs_curso/stoq/pkg/crypto"
)

func main() {

	dados, _ := crypto.ReadCryptoFile("./dados.txt")

	fmt.Println(string(dados))
}
