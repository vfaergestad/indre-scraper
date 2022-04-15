package links_db

import (
	"indre-scraper/internal/db"
	"log"
	"strings"
	"time"
)

const collection = "links"

// IsLinkInDB checks if link is in DB
func IsLinkInDB(link string) (bool, int) {
	doc, err := db.GetClient().Collection(collection).Doc(getLinkName(link)).Get(db.GetContext())
	if err != nil {
		return false, -1
	} else {
		ageDate, err := time.Parse("2006-01-02", doc.Data()["date"].(string))
		if err != nil {
			log.Println(err)
		}
		age := int(time.Since(ageDate).Hours() / 24)
		log.Println("Found duplicate link")
		return true, age
	}
}

// AddLinksToDB adds links to DB
func AddLinksToDB(links []string) error {
	for _, link := range links {
		err := addLinkToDB(link)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}

func addLinkToDB(link string) error {
	_, err := db.GetClient().Collection(collection).Doc(getLinkName(link)).Set(db.GetContext(), map[string]interface{}{
		"link": link,
		"date": time.Now().Format("2006-01-02"),
	})
	if err != nil {
		return err
	}
	return nil
}

func getLinkName(link string) string {
	splitLink := strings.Split(link, "/")
	return splitLink[1]
}
