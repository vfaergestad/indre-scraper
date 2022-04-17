package scraper

import (
	"indre-scraper/internal/db/articles_db"
	"log"
)

const (
	mainDomain = "www.indre.no"
	baseURL    = "https://www.indre.no"
)

func DoScrape() {
	log.Println("Scraping started")
	log.Println("Retrieving new links")
	links, err := GetLinks(baseURL, mainDomain)
	if err != nil {
		log.Fatal(err)
	}
	if len(links) == 0 {
		log.Println("No new links found")
		return
	} else {
		log.Printf("Found %d new links", len(links))
	}

	log.Println("Scraping links")
	for i, link := range links {
		log.Printf("%d of %d: %s", i+1, len(links), link)
		article, err := ScrapeArticle(link)
		if err != nil {
			log.Println(err)
		}
		err = articles_db.AddArticle(article)
		if err != nil {
			log.Println(err)
		}
	}
	log.Println("Scraping finished")
}
