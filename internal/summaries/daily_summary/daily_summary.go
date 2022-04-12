package daily_summary

import (
	"indre-scraper/internal/db/tags_db"
	"indre-scraper/internal/summaries/util"
	"strconv"
)

func GetDailySummaryLocation(date string) (string, error) {
	tags, err := tags_db.GetTagsFromDate(date)
	if err != nil {
		return "", err
	}
	for tag, _ := range tags {
		if !util.IsValidLocation(tag) {
			delete(tags, tag)
		}
	}

	sortedTags := util.SortTags(tags)

	var result string

	result += "Daily Summary:\n\n"
	result += "Aurskog-Høland: " + strconv.Itoa(util.AmountInAH(tags)) + "\n"
	result += "Lillestrøm: " + strconv.Itoa(util.AmountInLS(tags)) + "\n\n"
	result += "Each Location:\n\n"
	for _, tag := range sortedTags {
		result += "" + tag + ": " + strconv.Itoa(tags[tag]) + " \n"
	}

	return result, nil
}
