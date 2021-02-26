package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type people struct {
	Number int `json:"number"`
}

type csb20 []struct {
	S20ADRNR   int64  `json:"S20_ADR_NR"`
	S20ARTNR   int64  `json:"S20_ART_NR"`
	S20EANNR   int64  `json:"S20_EAN_NR"`
	S20EINHEIT string `json:"S20_EINHEIT"`
}

func main() {

	url := "http://192.168.71.64:20880/read20"

	csbClient := http.Client{
		Timeout: time.Second * 10, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "csb20")

	res, getErr := csbClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	csb20 := csb20{}
	jsonErr := json.Unmarshal(body, &csb20)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(csb20)
}
