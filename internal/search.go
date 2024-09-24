package internal

import "errors"

func (cont *Container) Search(id int) (article Article, err error) {
	if len(cont.articles) == 0 {
		return Article{}, errors.New("empty container")
	}
	result, ok := cont.articles[id]
	if !ok {
		return Article{}, errors.New("article not found")
	}
	return result, nil
}
