package parser

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"path/filepath"
	"ppeua/FRead/pkg"
)

// ParesUrl todo:将 #话题 换成对应的链接
/*
解析指定的url，并将图片+内容写入当前的目录下的title.md
*/
func ParesUrl(url string, markdownPath string) ([]string, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		//todo:构造一个标准的error输出格式
		return nil, errors.New("http get error: " + err.Error())
	}
	// 用goquery解析
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var title string
	doc.Find("div#detail-title.title").Each(func(i int, s *goquery.Selection) {
		title = s.Text()
	})
	var content string
	doc.Find("div#detail-desc.desc").Each(func(i int, s *goquery.Selection) {
		content = s.Text()
	})
	if len(title) == 0 {
		//按照汉字来处理 UTF-8格式
		var runeTitle []rune
		runeContent := []rune(content)
		if len(runeContent) <= 10 {
			runeTitle = runeContent
		} else {
			runeTitle = runeContent[:10]
		}
		title = string(runeTitle)
	}
	path := filepath.Join(markdownPath, title+".md")
	var imgUrl []string
	doc.Find("meta[name='og:image']").Each(func(i int, s *goquery.Selection) {
		img := s.AttrOr("content", "")
		imgUrl = append(imgUrl, img)
	})
	return pkg.Text2md(path, title, content, imgUrl...), nil
}
