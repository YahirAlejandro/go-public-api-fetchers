package telebotclient

import (
	"io/ioutil"
	"log"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func GetToken() string {
	tok, err := ioutil.ReadFile("../conf/token")
	check(err)

	return string(tok)
}
