package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"indre-scraper/internal/structs"
	"log"
	"net/http"
	"strings"
	"time"
)

func ScrapeArticle(link string) (structs.Article, error) {
	var article structs.Article
	article.Link = link
	url := baseURL + link
	res, err := http.Get(url)
	if err != nil {
		return structs.Article{}, err
	}

	if res.StatusCode != http.StatusOK {
		log.Printf("Error: %s", res.Status)
		return structs.Article{}, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return structs.Article{}, err
	}

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if s.AttrOr("property", "") == "article:tag" {
			tag := s.AttrOr("content", "")
			article.Tags = append(article.Tags, tag)
		}
		if s.AttrOr("name", "") == "author" {
			author := s.AttrOr("content", "")
			article.Author = author
		}
		if s.AttrOr("property", "") == "article:published_time" {
			publishedTime := s.AttrOr("content", "")
			publishedTime = strings.ReplaceAll(publishedTime, "+0200", "")
			publishedTime = strings.ReplaceAll(publishedTime, "+0100", "")
			parsed, err := time.Parse("2006-01-02T15:04:05.000", publishedTime)
			if err != nil {
				log.Printf("Error: %s", err)
				return
			}
			article.PublishedTime = parsed
		}
		if s.AttrOr("property", "") == "lp:premium" {
			article.Premium = s.AttrOr("content", "") == "true"
		}
	})
	return article, nil
}
