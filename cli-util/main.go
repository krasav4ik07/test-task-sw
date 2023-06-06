package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"log"
	"net/http"
	"regexp"
)

type RequestString struct {
	Str string `json:"str"`
}

func main() {
	str, url := GetParametersStrAndUrl()
	VerifyStrForSpaces(str)
	supstr := GetSupstring(str, url)
	log.Println("string -", *str, " supstring -", supstr)
}

func GetParametersStrAndUrl() (*string, *string) {
	str := flag.String("str", "-1",
		"Your string, example:\"abcabcbb\"")
	url := flag.String("url", "-1",
		"Url http server, example:\"http://localhost:8001/api/substring\"")
	flag.Parse()
	if *str == "-1" || *url == "-1" {
		log.Fatal(errors.New("parameters \"-str\" or \"-url\" not passed"))
	}
	log.Println("str -", *str)
	log.Println("url -", *url)
	return str, url
}

func GetSupstring(str *string, url *string) string {
	var supstring string
	var body bytes.Buffer
	err1 := json.NewEncoder(&body).Encode(RequestString{Str: *str})
	if err1 != nil {
		log.Fatal("data encoding error: ", err1)
	}

	response, err2 := http.Post(*url, "application/json", &body)
	if err2 != nil {
		log.Fatal("error request: ", err2)
	}

	if response.StatusCode == 404 {
		log.Fatal("error request: 404 page not found")
	}

	err3 := json.NewDecoder(response.Body).Decode(&supstring)
	if err3 != nil {
		log.Fatal("data decoding error: ", err3)
	}

	return supstring
}

func VerifyStrForSpaces(str *string) {
	matched, _ := regexp.MatchString(`^\S*$`, *str)
	if !matched {
		log.Fatal("string contains spaces")
	}
	return
}
