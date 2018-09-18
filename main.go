package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/**
* As WIP as physically possible
*
* This little app boy is going to eventually search through all of Craigslist instead of just limiting
* searches to one baby-back little bitch area. I'm sick of having to do the same searches in Philadelphia
* and South Jersey. Shit's annoying for sure. Why wouldn't they just allow multiple area searches? Who
* knows. Anywhere, here we go fellas.
*
* This baby has lots of info on the stuff we'll need to make the urls
* https://www.craigslist.org/about/bulk_posting_interface
*
* This guy is the only html parser I've looked at so far
* https://godoc.org/golang.org/x/net/html
 */

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
