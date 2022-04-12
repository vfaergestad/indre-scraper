package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"indre-scraper/internal/db/tags_db"
	"indre-scraper/internal/scraper/scrape_links"
	"log"
	"net/http"
)

const (
	mainDomain = "www.indre.no"
	baseURL    = "https://www.indre.no"
)

func InitScrape() {
	log.Println("Scraping started")
	links, err := scrape_links.GetLinks(baseURL, mainDomain)
	if err != nil {
		log.Fatal(err)
	}

	tags := countTags(links)
	err = tags_db.AddTodayTags(tags)
	if err != nil {
		log.Fatal(err)
	}
}

func countTags(links []string) map[string]int {
	tags := make(map[string]int)
	for _, link := range links {
		url := baseURL + link
		res, err := http.Get(url)
		if err != nil {
			log.Println(err)
		}

		if res.StatusCode != http.StatusOK {
			log.Printf("Error: %s", res.Status)
			continue
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		doc.Find("meta").Each(func(i int, s *goquery.Selection) {
			if s.AttrOr("property", "") == "article:tag" {
				tag := s.AttrOr("content", "")
				tags[tag]++
			}
		})
	}
	return tags
}
