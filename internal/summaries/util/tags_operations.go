package util

import "indre-scraper/internal/structs"

func CountTags(articles []structs.Article) map[string]int {
	tags := make(map[string]int)
	for _, article := range articles {
		for _, tag := range article.Tags {
			tags[tag]++
		}
	}
	return tags
}

// SortTags sorts the tags by their count
func SortTags(tags map[string]int) []string {
	var keys []string

	for key := range tags {
		keys = append(keys, key)
	}

	for i := 0; i < len(keys); i++ {
		for j := 0; j < len(keys)-i-1; j++ {
			if tags[keys[j]] < tags[keys[j+1]] {
				keys[j], keys[j+1] = keys[j+1], keys[j]
			}
		}
	}

	return keys
}
