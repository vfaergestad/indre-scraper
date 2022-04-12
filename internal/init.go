package internal

import (
	"indre-scraper/internal/db"
	"indre-scraper/internal/summaries/daily_summary"
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

	//scraper.InitScrape()

	summary, err := daily_summary.GetDailySummaryLocation(time.Now().Format("2006-01-02"))
	if err != nil {
		panic(err)
	}
	log.Println(summary)
}
