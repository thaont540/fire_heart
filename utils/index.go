package utils

import (
	"bytes"
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
)

func Env(key string) string {
	err := godotenv.Load()
	if err != nil {
		// for heroku
		return os.Getenv(key)
		//log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}

func AuthenticatedHttpGet(endpoint string) *http.Response {
	GetToken()
	client := &http.Client{}
	r, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer " + Ms365Token)

	res, err := client.Do(r)

	return res
}

func AuthenticatedHttpPost(endpoint string, data map[string]string) *http.Response {
	GetToken()
	dataBody, err := json.Marshal(data)

	client := &http.Client{}
	r, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(dataBody))
	if err != nil {
		log.Fatal(err)
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Content-Length", strconv.Itoa(len(data)))
	r.Header.Add("Authorization", "Bearer " + Ms365Token)

	res, err := client.Do(r)

	return res
}
