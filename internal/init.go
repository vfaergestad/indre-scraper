package internal

import (
	"indre-scraper/internal/db"
	"indre-scraper/internal/scraper"
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

	scraper.DoScrape()

}
