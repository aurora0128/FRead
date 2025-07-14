package global

import (
	"github.com/goccy/go-json"
	"log"
	"ppeua/FRead/model"
	"sync"
)

// 暴露给全局用
var (
	Once sync.Once
	Repo *model.ArticleRepo
)

// InitGlobal 为了保持干净防止循环依赖,初始化的函数手动传入
func InitGlobal(readrepo func() ([]byte, error)) {
	//初始化仓库
	Repo = &model.ArticleRepo{
		Size:       0,
		Mutex:      sync.RWMutex{},
		ArticleMap: make(map[string]*model.Article),
	}
	jsonData, err := readrepo()
	if err != nil {
		log.Fatal(err)
	}
	if len(jsonData) == 0 {
		log.Println("文件为空，跳过解析")
		return
	}

	Repo.Mutex.Lock()
	defer Repo.Mutex.Unlock()
	if err := json.Unmarshal(jsonData, Repo); err != nil {
		log.Fatalln("32 ", err)
	}
}
