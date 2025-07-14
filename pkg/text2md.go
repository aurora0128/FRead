package pkg

import (
	"errors"
	"fmt"
	"github.com/gofrs/uuid/v5"
	"log"
	"os"
	"path/filepath"
)

/*
返回列表：
0:title
1:content
2:第一张图片
*/
func Text2md(path, title, content string, img ...string) []string {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		var pathErr *os.PathError
		if !errors.As(err, &pathErr) {
			log.Println(err)
			return nil
		}
		
		//todo:合适的命名规则
		path = filepath.Dir(path)
		id, _ := uuid.NewV4()
		fmt.Printf("名字错误发生替换:%s -> %s\n", filepath.Base(path), id.String()[:6])
		path = filepath.Join(path, id.String()[:6]+".md")

		return Text2md(path, title, content, img...)
	}
	defer f.Close()
	for _, url := range img {
		content = content + "![" + "](" + url + ")\n"
	}
	f.Write([]byte(title + content + "\n"))
	return []string{title, content, img[0]}
}
