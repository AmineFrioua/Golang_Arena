package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	passwordLength = 16
	passwordChars  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	password := make([]byte, passwordLength)
	for i := range password {
		password[i] = passwordChars[rand.Intn(len(passwordChars))]
	}

	fmt.Println(string(password))
}