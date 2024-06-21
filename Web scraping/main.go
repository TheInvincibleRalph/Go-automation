package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Step 1: Make an HTTP request

	response, err := http.Get("https://scaler.com")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Step 2: Parse the HTML

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Step 3: Find and extract the data

	document.Find("h1").Each(func(index int, element *goquery.Selection) { //Iterates over each h1 element found.
		// Get the text of the h1 element
		heading := element.Text()
		fmt.Println("Heading:", heading)
	})
}
