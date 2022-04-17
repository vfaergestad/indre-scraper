package articles_db

import (
	"google.golang.org/api/iterator"
	"indre-scraper/internal/db"
	"indre-scraper/internal/structs"
	"strings"
	"time"
)

const collection = "articles"

func AddArticle(article structs.Article) error {
	_, err := db.GetClient().Collection(collection).Doc(getIDFromLink(article.Link)).Set(db.GetContext(), article)
	if err != nil {
		return err
	}
	return nil
}

func GetArticle(link string) (structs.Article, error) {
	article := structs.Article{}
	doc, err := db.GetClient().Collection(collection).Doc(getIDFromLink(link)).Get(db.GetContext())
	if err != nil {
		return structs.Article{}, err
	}
	err = doc.DataTo(&article)
	if err != nil {
		return structs.Article{}, err
	}
	return article, nil
}

func GetArticlesFromRange(startDate time.Time, endDate time.Time) ([]structs.Article, error) {
	var articles []structs.Article
	query := db.GetClient().Collection(collection).Where("PublishedTime", ">=", startDate).Where("PublishedTime", "<=", endDate)
	iter := query.Documents(db.GetContext())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var article structs.Article
		err = doc.DataTo(&article)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func getIDFromLink(link string) string {
	splitLink := strings.Split(link, "/")
	return splitLink[len(splitLink)-1]
}
