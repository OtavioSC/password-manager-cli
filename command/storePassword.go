package command

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"log"
)

func StorePassword(Args []string) error {
	a := Args[0]
	storedPass, err := EncryptPassword(a)
	if err != nil {
		log.Fatal("Failed to store password")
	}
	fmt.Printf("%s", storedPass)
	return nil
}

func EncryptPassword(password string) (string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := &privateKey.PublicKey
	m := md5.New()
	b := []byte(password)
	m.Write(b)
	encryptedPass, err := rsa.EncryptOAEP(m, rand.Reader, publicKey, b, nil)
	return base64.StdEncoding.EncodeToString(encryptedPass), err
}
