package service

import (
	"ppeua/FRead/internal/global"
	"ppeua/FRead/model"
	"sort"
)

//func GetArticlesOld() (*[]model.Article, error) {
//	return &global.Repo.Articles, nil
//}

func GetArticles() (*[]model.Article, error) {
	//对hash遍历 并按照添加时间逆序
	var articles []model.Article
	for _, article := range global.Repo.ArticleMap {
		articles = append(articles, *article)
	}

	//按照修改时间降序来排列(从新到旧)
	sort.Slice(articles, func(i, j int) bool {
		return articles[i].UpdatedAt.After(articles[j].UpdatedAt)
	})
	return &articles, nil
}
