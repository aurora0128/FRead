package service

import (
	"errors"
	"ppeua/FRead/internal/global"
	"ppeua/FRead/internal/repo"
)

func DeleteArticle(id string) error {
	global.Repo.Mutex.Lock()
	defer global.Repo.Mutex.Unlock()
	if _, exit := global.Repo.ArticleMap[id]; exit {
		global.Repo.Size--
		delete(global.Repo.ArticleMap, id)
		//写入repo
		repo.WriteArticles()
	} else {
		return errors.New("article not exist")
	}
	return nil
}
