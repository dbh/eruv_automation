package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"io"

	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	parshaUrl := os.Getenv("PARSHA_URL")
	log.Printf("parsha_url: %s\n", parshaUrl)

	response, err := http.Get(parshaUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	content, err := io.ReadAll(io.Reader(response.Body))
	if err != nil {
		log.Fatal(err)
	}

	// TODO: parse the content
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(content)))
	if err != nil {
		log.Fatal(err)
	}
	// var h2 string
	var parsha string
	doc.Find("h2").First().Each(func(i int, s *goquery.Selection) {
		parsha = s.Children().Nodes[0].NextSibling.Data
		parsha = strings.TrimSpace(parsha)
		log.Printf(("parsha: %s\n"), parsha)
	})
	// TODO: update the eruv email
	// TODO: send the email

}
