package controller

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/mikecrinite/craigslist-global/model"
)

var scheme = "https://"                        // Craiglist uses HTTPS protocol
var base = ".craigslist.org"                   // URL base for the Craigslist domain
var search = "/search/"                        // Currently, this will always immediately follow the base
var query = "?query="                          // Denotes the queryString, required by craigslist in order to complete the search
var imageCL = "https://images.craigslist.org/" // CL's image hosting site
var imageTail = "_600x450.jpg"                 // Contains the resolution and filetype of the image

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

// Post contains all the pertinent information to make a Craiglist post
type Post struct {
	PostLink string // The URL of the post
	DataIds  string // The value of the data-ids field (which is a collection of image ids)
	Title    string // The title of the post
	Price    string // The price of the vehicle listed by the owner
	Region   string // The region of the post
	HasImage bool   // Does the listing have an image
}

// Currently returns the first image id as a link to the image instead of the weird ID format CL uses
func prep(d string) (string, bool) {
	l := strings.Split(d, ",") // This will be the list of ids, of which we only want the first
	if len(l) > 0 && l[0] != "" {
		id := strings.Split(l[0], ":")[1] // This gives the id itself
		return imageCL + id + imageTail, true
	}
	return "", false
}

// ScrapeCL scrapes craigslist searches for the given category and keywords across all regions
func ScrapeCL(category string, keywords string) []Post { //[]string {
	if &keywords == nil || keywords == "" {
		return []Post{}
	}
	linkMap := make(map[string]Post)
	linkChannel := make(chan Post)
	doneChannel := make(chan bool)

	go func() {
		for {
			p, more := <-linkChannel
			linkMap[p.PostLink] = p
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

	c.OnHTML("a.result-image.gallery", func(h *colly.HTMLElement) {
		href := h.Attr("href")
		data := h.Attr("data-ids")
		title := ""
		region := ""
		price := h.ChildText("span")

		linkChannel <- Post{href, data, title, price, region, false}
	})

	c.OnHTML("a.button.next", func(h *colly.HTMLElement) {
		u := h.Attr("href")
		if &u != nil && u != "" {
			nextURL := strings.Split(h.Response.Request.URL.String(), "/search")[0] + u
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
	links := make([]Post, len(linkMap))
	i := 0
	for _, post := range linkMap {
		title, region := stripMetaFromURL(post.PostLink)
		imageLink, hasImg := prep(post.DataIds)
		post.DataIds = imageLink // Get a display image
		post.HasImage = hasImg
		post.Region = region
		post.Title = title
		links[i] = post
		i++
	}

	return links
}

// CleanForQuery removes separating spaces and replaces them with a '+'
func CleanForQuery(dirty string) string {
	return strings.Replace(dirty, " ", "+", -1)
}

func stripMetaFromURL(url string) (string, string) {
	var region string
	l := strings.Split(url, "/")
	r := strings.Split(url, ".")
	if len(r) > 2 {
		region = strings.Replace(r[0], "https://", "", 1)
	}
	if len(l) > 2 {
		dirtyTitle := l[len(l)-2]
		return strings.Replace(dirtyTitle, "-", " ", -1), region
	}

	return url, region
}
