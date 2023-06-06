package controllers

import (
	"encoding/json"
	"errors"
	"http-server/internal/usecases"
	"log"
	"net/http"
)

type RequestString struct {
	Str string `json:"str"`
}

type SubstringController struct {
	substring *usecases.SubstringUsecases
}

func NewSubstringController(s *usecases.SubstringUsecases) *SubstringController {
	return &SubstringController{substring: s}
}

func (c *SubstringController) GetString(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var data RequestString
	err1 := json.NewDecoder(req.Body).Decode(&data)
	if err1 != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(err1)
		log.Println(err1)
		return
	}

	if data.Str == "" {
		err2 := errors.New("empty string")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(err2)
		log.Println(err2)
		return
	}

	substring, err3 := c.substring.SearchMaxUniqueSubstring(data.Str)
	if err3 != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(err3)
		log.Println(err3)
		return
	}

	log.Println("String -", data.Str, len([]rune(data.Str)))
	log.Println("SubString -", substring, len([]rune(substring)))

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(substring)
}
