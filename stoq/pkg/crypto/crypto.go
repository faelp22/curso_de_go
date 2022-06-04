package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

func encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

func decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return data, nil
}

func WriteCryptoFile(filename string, data []byte) (ok bool) {

	key := []byte(`6A77BC83633EE51864C4721EC34BB6A7`)

	ciphertext, err := encrypt(key, data)
	if err != nil {
		log.Println(err.Error())
		log.Println("WriteCryptoFile, Erro ao encriptar arquivo.")
		return false
	}

	err = ioutil.WriteFile(filename, ciphertext, 0600)
	if err != nil {
		log.Println("WriteCryptoFile, Erro ao gravar arquivo.")
		return false
	}
	return true
}

func ReadCryptoFile(filename string) (data []byte, ok bool) {

	key := []byte(`6A77BC83633EE51864C4721EC34BB6A7`)

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err.Error())
		log.Println("ReadCryptoFile, Erro ao ler arquivo.")
		return data, false
	}

	data, err = decrypt(key, bytes)
	if err != nil {
		log.Println("ReadCryptoFile, Erro ao desencriptar arquivo.")
		return data, false
	}

	return data, true
}
