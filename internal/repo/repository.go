package repo

import (
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"ppeua/FRead/internal/config"
	"ppeua/FRead/internal/global"
	"ppeua/FRead/model"

	"github.com/goccy/go-json"
	"github.com/gofrs/uuid/v5"
)

/*
filename:repository_size
*/
func getRepoPath() string {
	//todo：repo文件名以及分配机制 需要重新考虑
	return filepath.Join(config.Cfg.Storage.MarkdownPath, "repo.json")
}

// todo:对repo访问需要考虑锁
// todo:优化一下输出的json格式(考量是否删去repo保留hash)
func WriteRepo(article *model.Article) error {
	file, err := os.OpenFile(getRepoPath(), os.O_WRONLY, 0666)
	if err != nil {
		return errors.New("repo write err: " + err.Error())
	}
	defer file.Close()

	global.Repo.Size++
	global.Repo.Mutex.Lock()
	global.Repo.ArticleMap[article.ID] = article
	global.Repo.Mutex.Unlock()

	jsonData, err := json.MarshalIndent(global.Repo, "", "\t")
	if err != nil {
		return errors.New("json err: " + err.Error())
	}
	file.Write(jsonData)
	return nil
}

func WriteArticles() error {
	file, err := os.OpenFile(getRepoPath(), os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return errors.New("repo writeArticles err: " + err.Error())
	}
	defer file.Close()
	jsonData, err := json.MarshalIndent(global.Repo, "", "\t")
	if err != nil {
		return errors.New("json err: " + err.Error())
	}
	file.Write(jsonData)
	return nil
}

func ReadRepo() ([]byte, error) {

	if _, err := os.Stat(getRepoPath()); os.IsNotExist(err) {
		// 创建空文件（不会写入任何内容）
		if err := os.WriteFile(getRepoPath(), []byte(""), 0666); err != nil {
			log.Fatalf("无法创建文件: %v", err)
		}
	}

	jsonData, err := os.ReadFile(getRepoPath())
	if err != nil {
		if err != io.EOF {
			return nil, errors.New("ReadRepo err: " + err.Error())
		}
	}

	return jsonData, nil
}

//将content以markdown格式保存到本地

func WriteMd2Repo(path, title, content string, img ...string) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, os.ModePerm)

	//发生了重复名字错误
	if err != nil {
		var pathErr *os.PathError
		if !errors.As(err, &pathErr) {
			log.Println(err)
			return
		}

		//todo:合适的命名规则
		path = filepath.Dir(path)
		id, _ := uuid.NewV4()
		log.Printf("名字错误发生替换:%s -> %s\n", filepath.Base(path), id.String()[:6])
		path = filepath.Join(path, id.String()[:6]+".md")
		WriteMd2Repo(path, title, content, img...)
	}

	defer f.Close()

	f.Write([]byte(title + content + "\n"))
}
