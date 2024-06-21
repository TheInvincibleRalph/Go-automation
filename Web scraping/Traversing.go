package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Fetch the HTML document
	response, err := http.Get("https://scaler.com")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Parse the HTML document
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body.", err)
	}

	// Find all <div> elements and their direct children
	document.Find("div").Children().Each(func(index int, element *goquery.Selection) {
		fmt.Println("Child of <div>:", element.Text())
	})

	// Find the parent of <span> elements
	document.Find("span").Parent().Each(func(index int, element *goquery.Selection) {
		fmt.Println("Parent of <span>:", element.Text())
	})

	// Find all siblings of <h1> elements
	document.Find("h1").Siblings().Each(func(index int, element *goquery.Selection) {
		fmt.Println("Sibling of <h1>:", element.Text())
	})

	// Find the next sibling of <h1> elements
	document.Find("h1").Next().Each(func(index int, element *goquery.Selection) {
		fmt.Println("Next Sibling of <h1>:", element.Text())
	})

	// Find the previous sibling of <h1> elements
	document.Find("h1").Prev().Each(func(index int, element *goquery.Selection) {
		fmt.Println("Previous Sibling of <h1>:", element.Text())
	})

	// Find all <p> elements that are descendants of <div> elements
	document.Find("div").Find("p").Each(func(index int, element *goquery.Selection) {
		fmt.Println("Descendant <p> of <div>:", element.Text())
	})
}
