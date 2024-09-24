package internal

import (
	"encoding/json"
	"errors"
)

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	ID      int    `json:"id"`
}

func ParseJSON(data []byte) (art Article, err error) {
	var article Article
	err = json.Unmarshal(data, &article)
	if err != nil {
		return Article{}, errors.New("invalid json syntax")
	}
	if article.ID == 0 || article.Title == "" || article.Content == "" {
		return Article{}, errors.New("invalid article data")
	}
	return article, nil
}

type Container struct {
	articles map[int]Article
}

func NewContainer() *Container {
	return &Container{
		articles: make(map[int]Article),
	}
}

func (cont *Container) AddArticle(article Article) (err error) {
	if article == (Article{}) {
		return errors.New("empty article")
	}
	cont.articles[article.ID] = article
	return nil
}

func (cont *Container) GetArticles() map[int]Article {
	return cont.articles
}
