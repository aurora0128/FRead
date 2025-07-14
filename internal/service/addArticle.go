package service

import (
	"github.com/gofrs/uuid/v5"
	"log"
	"ppeua/FRead/internal/config"
	"ppeua/FRead/internal/parser"
	"ppeua/FRead/internal/repo"
	"ppeua/FRead/model"
	"time"
)

func AddArticle(url string) (*model.Article, error) {
	url, err := parser.Raw2Url(url)
	if err != nil {
		log.Printf("Raw2Url:%s,err: %s", url, err.Error())
		return nil, err
	}
	info, err := parser.ParesUrl(url, config.Cfg.Storage.MarkdownPath)
	if err != nil {
		log.Printf("ParesUrl:%s,err: %s", url, err.Error())
		return nil, err
	}

	id, err := uuid.NewV4()
	if err != nil {
		log.Printf("UUID:%s,err: %s", id, err.Error())
		return nil, err
	}

	article := &model.Article{
		ID:        id.String(), //test
		URL:       url,
		Title:     info[0],
		Content:   info[1],
		Thumbnail: info[2],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := repo.WriteRepo(article); err != nil {
		log.Printf("WriteRepo:%s,err: %s", article, err.Error())
		return nil, err
	}
	return article, nil
}
