package controller

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/mikecrinite/craigslist-go/model"
)

var scheme = "https://"      // Craiglist uses HTTPS protocol
var base = ".craigslist.org" // URL base for the Craigslist domain
var search = "/search/"      // Currently, this will always immediately follow the base
var query = "?query="        // Denotes the queryString, required by craigslist in order to complete the search

// BuildURL creates a craigslist URL from region, category, and keywords provided
func buildURL(region string, category string, keywords string) string {
	if &region == nil {
		log.Fatal("Region cannot be nil")
	}
	if &category == nil {
		log.Fatal("Category cannot be nil")
	}
	if &keywords == nil {
		log.Fatal("Keywords cannot be nil")
	}
	url := scheme + region + base + search + category + query + keywords
	return url
}

// ScrapeCL scrapes craigslist searches for the given category and keywords across all regions
func ScrapeCL(category string, keywords string) []string {
	if &keywords == nil || keywords == "" {
		return []string{}
	}

	linkSet := make(map[string]struct{})
	linkChannel := make(chan string)
	doneChannel := make(chan bool)

	go func() {
		for {
			t, more := <-linkChannel
			linkSet[t] = struct{}{}
			if !more {
				break
			}
		}

		doneChannel <- true
		return
	}()

	c := colly.NewCollector(
		colly.Async(true),
	)

	c.URLFilters = []*regexp.Regexp{
		regexp.MustCompile(`^https?://[a-z]+\.craigslist\.org/*`),
	}

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Making request to: ", r.URL)
	})

	c.OnHTML("a.result-title.hdrlnk", func(h *colly.HTMLElement) {
		t := h.Attr("href")
		linkChannel <- t
	})

	c.OnHTML("a.button.next", func(h *colly.HTMLElement) {
		u := h.Attr("href")
		if &u != nil && u != "" {
			nextURL := strings.Split(h.Response.Request.URL.String(), "/search")[0] + u
			//fmt.Println("Followed Next: ", nextURL)
			c.Visit(nextURL)
		}
	})

	c.Limit(&colly.LimitRule{
		Parallelism: 3,
		RandomDelay: 3 * time.Second,
	})

	for _, region := range model.Regions {
		url := buildURL(region, category, keywords)
		err := c.Visit(url)
		if err != nil {
			log.Fatal(err)
		}
	}

	c.Wait()
	close(linkChannel)

	<-doneChannel

	// Using append() is about 20% slower than directly assigning values
	links := make([]string, len(linkSet))
	i := 0
	for l := range linkSet {
		//fmt.Println(strconv.Itoa(i) + " - " + l)
		links[i] = l
		i++
	}

	return links
}

// CleanForQuery removes separating spaces and replaces them with a '+'
func CleanForQuery(dirty string) string {
	return strings.Replace(dirty, " ", "+", -1)
}
