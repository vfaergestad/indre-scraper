package scraper

import (
	"github.com/gocolly/colly"
	"indre-scraper/internal/db/articles_db"
	"strings"
)

func GetLinks(url string, domain string) ([]string, error) {
	c := colly.NewCollector(
		colly.AllowedDomains(domain),
	)

	var links []string

	c.OnHTML(".maelstrom-wrapper", func(e *colly.HTMLElement) {
		links = e.ChildAttrs("a[href]", "href")
		links = cleanLinks(links)
	})

	err := c.Visit(url)
	if err != nil {
		return nil, err
	}

	return links, nil
}

func cleanLinks(links []string) []string {
	var cleanedLinks []string
	for _, link := range links {
		if link != "" && !strings.Contains(link, "https") && !strings.Contains(link, ".ece") {
			if _, err := articles_db.GetArticle(link); err != nil {
				cleanedLinks = append(cleanedLinks, link)
			}
		}
	}
	cleanedLinks = removeDuplicates(cleanedLinks)
	return cleanedLinks
}

func removeDuplicates(strSlice []string) []string {
	allKeys := make(map[string]bool)
	var list []string
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
