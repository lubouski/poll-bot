package main

import (
	"fmt"
	"encoding/json"
	"log"
	"bytes"
	"net/http"
)

type resultUrl struct {
	Result_url string
}

func main() {
	url := "https://dev.by"
	fmt.Println(urlShort("url=" + url))
}

func urlShort(url string) string {
	var urlStr = []byte(url)
	response, err := http.Post("https://cleanuri.com/api/v1/shorten", "application/x-www-form-urlencoded",
        bytes.NewBuffer(urlStr))

    	if err != nil {
        	log.Fatal(err)
    	}

    	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	var res resultUrl
	err = decoder.Decode(&res)

	if err != nil {
		panic(err)
	}

	return res.Result_url
}
