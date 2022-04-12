package tags_db

import (
	"indre-scraper/internal/db"
	"time"
)

const collection = "tags"

// AddTodayTags adds the given tags to the database
func AddTodayTags(tags map[string]int) error {
	_, err := db.GetClient().Collection(collection).Doc(time.Now().Format("2006-01-02")).Set(db.GetContext(), tags)
	if err != nil {
		return err
	}
	return nil
}

func GetTagsFromDate(date string) (map[string]int, error) {
	var tags map[string]int
	doc, err := db.GetClient().Collection(collection).Doc(date).Get(db.GetContext())
	if err != nil {
		return tags, err
	}
	err = doc.DataTo(&tags)
	if err != nil {
		return tags, err
	}

	return tags, nil
}
