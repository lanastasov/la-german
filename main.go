package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func scrape(url string) {
	// Fetch the URL
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching URL: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", res.StatusCode, res.Status)
	}

	// Parse the HTML
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalf("Error parsing HTML: %v", err)
	}

	// Find the desired element and print its text
	text := doc.Find(".col.text").First().Text()
	fmt.Println(text)
}

func main() {
	url := "https://app.memrise.com/community/course/46/1000-words-of-elementary-german/1/" // Replace with your target URL
	scrape(url)
}
