package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println(byteArrayAsString(executeRequest("https://philadelphia.craigslist.org/search/cto?query=subaru+impreza")))
}

func executeRequest(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err) // TODO: This, but better
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err) // TODO: This, but better
	}
	return body
}

func byteArrayAsString(arr []byte) string {
	return string(arr[:])
}
