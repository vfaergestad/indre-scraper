package internal

import (
	"indre-scraper/internal/db"
	"indre-scraper/internal/db/articles_db"
	"indre-scraper/internal/summaries/util"
	"log"
	"time"
)

func Init() {
	// Initialize the database.
	err := db.InitializeFirestore()
	if err != nil {
		panic(err)
	}

	defer func() {
		err = db.CloseFirestore()
		if err != nil {
			panic(err)
		}
	}()

	//scraper.DoScrape()

	articles, err := articles_db.GetArticlesFromRange(time.Now().Add(-24*30*time.Hour), time.Now())
	if err != nil {
		panic(err)
	}

	tags := util.CountTags(articles)
	sortedTags := util.SortTags(tags)
	for _, tag := range sortedTags {
		log.Printf("%s: %d", tag, tags[tag])
	}
}
