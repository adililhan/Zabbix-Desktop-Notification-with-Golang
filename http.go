package main

import (
	"net/http"
	"bytes"
	"io/ioutil"
)

func GetContent(url, json string) []uint8 {
	jsonReady := []byte(json)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonReady))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	defer response.Body.Close()

	check(err)

	body, err := ioutil.ReadAll(response.Body)

	return body
}