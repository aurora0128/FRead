package service

import (
	"errors"
	"ppeua/FRead/internal/global"
	"ppeua/FRead/model"
)

func GetArticle(id string) (*model.Article, error) {
	//调用map进行类型匹配
	if article, exit := global.Repo.ArticleMap[id]; exit {
		return article, nil
	} else {
		return article, errors.New("article not found")
	}
}
