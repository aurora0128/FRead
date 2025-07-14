package model

import "sync"

type ArticleRepoOld struct {
	Size         int       `json:"size"`
	Articles     []Article `json:"articles"`
	ArticlesHash ArticleList
}

type ArticleRepo struct {
	Size       int                 `json:"size"`
	Mutex      sync.RWMutex        `json:"-"`
	ArticleMap map[string]*Article `json:"articleHash"`
}
