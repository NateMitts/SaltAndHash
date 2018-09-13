package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func getPWD() []byte {
	fmt.Println("Enter a password")
	var pwd string
	_, err := fmt.Scan(&pwd)
	if err != nil {
		log.Println(err)

	}
	return []byte(pwd)
}

func hashAndSalt(pwd []byte) string {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {

		log.Println(err)

	}

	return string(hash)

}
func main() {
	for {
		pwd := getPWD()
		hash := hashAndSalt(pwd)
		fmt.Println("Salted Hash: ", hash)

	}

}
