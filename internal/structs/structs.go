package structs

import "time"

type Article struct {
	Link          string
	Author        string
	PublishedTime time.Time
	Tags          []string
	Premium       bool
}
