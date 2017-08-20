package config

import (
	"fmt"
	"io/ioutil"
)

// Public Variable
var (
	Token string
)

// ReadConfig Returns the token to use in the API
func ReadConfig() (string, error) {
	fmt.Println("Reading from API key file...")

	file, err := ioutil.ReadFile("./apikey.txt")
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	fmt.Println("Success!")
	Token = string(file)
	return Token, nil
}
