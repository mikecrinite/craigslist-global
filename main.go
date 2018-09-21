package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

// Run the application with currently-hard-coded values
func main() {
	//fmt.Println(url)
	//links := scrapeCL(categoryMap["cars & trucks - by owner"], "subaru")
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":      "craigslist-global",
			"categories": categoryMapKeys(),
		})
	})
	r.POST("/", func(c *gin.Context) {
		query := c.PostForm("search")
		category := c.PostForm("category")

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":      "craigslist-global",
			"categories": categoryMapKeys(),
			"links":      scrapeCL(categoryMap[category], cleanForQuery(query)),
		})
	})
	r.Run(":8095") // 0.0.0.0:8095
}

func cleanForQuery(dirty string) string {
	return strings.Replace(dirty, " ", "+", -1)
}

var scheme = "https://"      // Craiglist uses HTTPS protocol
var base = ".craigslist.org" // URL base for the Craigslist domain
var search = "/search/"      // Currently, this will always immediately follow the base
var query = "?query="        // Denotes the queryString, required by craigslist in order to complete the search

// Builds a URL from region, category, and keywords provided
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

// Scrapes craigslist searches for the given category and keywords across all regions
func scrapeCL(category string, keywords string) []string {
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

	for _, region := range regions {
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

// Executes an http GET request for the provided URL, and if successful, returns the web page contents as a byte array
func executeRequest(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

// Returns the string representation of the provided byte array
func byteArrayAsString(arr []byte) string {
	return string(arr[:])
}

// Gets a full slice of keys from the categoryMap below
func categoryMapKeys() []string {
	keys := make([]string, len(categoryMap))

	i := 0
	for k := range categoryMap {
		keys[i] = k
		i++
	}

	return keys
}

// A map of human-readable categories to their Craigslist shorthand. This map is not currently a comprehensive mapping of the full list of categories
// This map is not currently immutable but should not be mutated anywhere in the codebase
var categoryMap = map[string]string{
	"apts broker fee":                     "fee",
	"apts broker no fee":                  "nfb",
	"apts/housing for rent":               "apa",
	"office & commercial":                 "off",
	"real estate - by broker":             "reb",
	"vacation rentals":                    "vac",
	"accounting/finance":                  "acc",
	"admin/office":                        "ofc",
	"architect/engineer/cad":              "egr",
	"art/media/design":                    "med",
	"business/mgmt":                       "bus",
	"customer service":                    "csr",
	"education/teaching":                  "edu",
	"et cetera":                           "etc",
	"food/beverage/hospitality":           "fbh",
	"general labor":                       "lab",
	"government":                          "gov",
	"healthcare":                          "hea",
	"human resource":                      "hum",
	"legal/paralegal":                     "lgl",
	"manufacturing":                       "mnu",
	"marketing/advertising/pr":            "mar",
	"nonprofit":                           "npo",
	"real estate":                         "rej",
	"retail/wholesale":                    "ret",
	"sales":                               "sls",
	"salon/spa/fitness":                   "spa",
	"science/biotech":                     "sci",
	"security":                            "sec",
	"skilled trades/artisan":              "trd",
	"software/qa/dba/etc":                 "sof",
	"systems/networking":                  "sad",
	"technical support":                   "tch",
	"transportation":                      "trp",
	"tv/film/video/radio":                 "tfr",
	"web/html/info design":                "web",
	"writing/editing":                     "wri",
	"appliances - by dealer":              "ppd",
	"atvs, utvs, snowmobiles - by dealer": "snd",
	"auto parts - by dealer":              "ptd",
	"auto wheels & tires - by dealer":     "wtd",
	"boats - by dealer":                   "bod",
	"business/commercial - by dealer":     "bfd",
	"cars & trucks - by owner/dealer":     "cta",
	"cars & trucks - by dealer":           "ctd",
	"cars & trucks - by owner":            "cto",
	"cell phones - by dealer":             "mod",
	"computers - by dealer":               "syd",
	"electronics - by dealer":             "eld",
	"farm & garden - by dealer":           "grq",
	"furniture - by dealer":               "fud",
	"general for sale - by dealer":        "fod",
	"heavy equipment - by dealer":         "hvd",
	"household items - by dealer":         "hsd",
	"materials - by dealer":               "mad",
	"motorcycles/scooters - by dealer":    "mcd",
	"rvs - by dealer":                     "rvd",
	"tickets - by dealer":                 "tid",
	"trailers - by dealer":                "trd",
}

// A list of all Craiglist regions for the application to loop through on its quest to get results from every region
var regions = []string{
	"albany",
	"allentown",
	"altoona",
	"annapolis",
	"baltimore",
	"binghamton",
	"catskills",
	"cnj",
	"charlottesville",
	"chambersburg",
	"delaware",
	"newlondon",
	"easternshore",
	"martinsburg",
	"elmira",
	"fingerlakes",
	"frederick",
	"fredericksburg",
	"harrisburg",
	"harrisonburg",
	"hartford",
	"hudsonvalley",
	"ithaca",
	"jerseyshore",
	"lancaster",
	"longisland",
	"newhaven",
	"newyork",
	"norfolk",
	"newjersey",
	"nwct",
	"oneonta",
	"philadelphia",
	"poconos",
	"reading",
	"richmond",
	"scranton",
	"southjersey",
	"smd",
	"pennstate",
	"syracuse",
	"utica",
	"washingtondc",
	"westernmass",
	"westmd",
	"williamsport",
	"winchester",
	"york",
}
