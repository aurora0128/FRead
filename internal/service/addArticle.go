package service

import (
	"log"
	"ppeua/FRead/internal/config"
	"ppeua/FRead/internal/parser"
	"ppeua/FRead/internal/repo"
	"ppeua/FRead/model"
	"ppeua/FRead/pkg"
	"strings"
	"time"

	"github.com/gofrs/uuid/v5"
)

// 目前支持列表:xhs zhihu
func AddArticle(url string) (*model.Article, error) {
	var err error
	url, err = pkg.Raw2Url(url)
	if err != nil {
		log.Printf("Raw2Url:%s,err: %s", url, err.Error())
		return nil, err
	}
	var info []string

	//如何判断谁是谁的链接?
	//1. 根据host:zhihu xiaohongshu
	if strings.Contains(url, config.Cfg.Platform.Xiaohongshu) {
		info, err = parser.ParesUrl(url, config.Cfg.Storage.MarkdownPath)
	} else if strings.Contains(url, config.Cfg.Platform.Zhihu) {
		info, err = parser.ParserUrlZhihu(url, config.Cfg.Storage.MarkdownPath)
	}

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
