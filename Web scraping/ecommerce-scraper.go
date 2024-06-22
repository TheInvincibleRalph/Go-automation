package main

import (
	"fmt"
	// "io"
	"encoding/csv"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

// custom type to keep scraped items
type Cars struct {
	name, price string
}

func main() {
	//download the targeted HTML document
	response, err := http.Get("https://jiji.ng/cars")
	if err != nil {
		log.Fatal("Failed to get targeted HTML page: ", err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		log.Fatalf("HTTP Error %d: %s", response.StatusCode, response.Status)
	}

	// fmt.Println(response.Body)

	// // Convert the response buffer to bytes

	// byteBody, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatal("Error while converting buffer", err)
	// }

	// // Convert the byte HTML to string and print it

	// stringBody := string(byteBody)
	// fmt.Println(stringBody)

	//parse the HTML document
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Failed to parse the HTML document: ", err)
	}

	//where to store the scraped data (products is a slice of Product struct)
	var products []Cars

	// bookHTMLElement := document.Find("div.b-list-advert__gallery__item").First()
	// fmt.Println(bookHTMLElement)

	// bookName := bookHTMLElement.Find("div.price").Text()
	// fmt.Println(bookName)

	//retrieve name and price from each product and print it
	document.Find("div.b-list-advert__gallery__item").Each(func(i int, p *goquery.Selection) {
		//scraping logic
		product := Cars{}
		product.name = p.Find("div.b-list-advert-base__data__title").Text()
		product.price = p.Find("div.b-list-advert-base__data__price").Text()

		//store the scraped items
		products = append(products, product)

	})

	fmt.Println(products)

	/*
	   In the .Eachfunc(i int, p *goquery.Selection) {...}) block, the ellipses is the part of the code
	   where you create an instance of the Car struct and extract the name and price of each element
	   and storing them in the struct instance after which you will append to the original struct.
	*/

	file, err := os.Create("products.csv")
	if err != nil {
		log.Fatal("Failed to create the output CSV file", err)
	}
	defer file.Close()

	//initialize a file writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	//define the CSV headers
	headers := []string{
		"name",
		"price",
	}

	//write the column headers
	writer.Write(headers)

	//add each product to the CSV file
	for _, product := range products {
		//convert a Product to an array of strings
		record := []string{
			product.name,
			product.name,
		}

		//write a new CSV record
		writer.Write(record)
	}

}
