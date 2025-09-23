package parser

import (
	"errors"
	"log"
	"ppeua/FRead/internal/config"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type zhihu struct {
	Url     string   `json:"url"`
	Title   string   `json:"question"`
	Content string   `json:"content"`
	Img     []string `json:"img"`
}

func cleanContent(content string) string {
	// 将多个换行压缩为一个换行
	content = regexp.MustCompile(`\r?\n+`).ReplaceAllString(content, "\n")
	// 将连续的空格和制表符压成一个空格（不影响换行）
	content = regexp.MustCompile(`[ \t]+`).ReplaceAllString(content, " ")
	content = strings.TrimSpace(content)
	return content
}
func ParserUrlZhihu(url string, markdownPath string) ([]string, error) {
	c := colly.NewCollector()
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*zhihu.com",
		Parallelism: 1,
		Delay:       2 * time.Second,
	})
	var result zhihu
	result.Img = make([]string, 1)
	result.Url = url
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		r.Headers.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
		r.Headers.Set("Connection", "keep-alive")

		// 添加Cookie
		r.Headers.Set("Cookie", config.Cfg.Cookie.ZhihuCookie)
	})
	c.OnHTML("h1.QuestionHeader-title", func(e *colly.HTMLElement) {
		result.Title = strings.TrimSpace(e.Text)
	})
	c.OnHTML("span.RichText", func(e *colly.HTMLElement) {
		result.Content = cleanContent(e.Text)
	})
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("请求失败: %v\n", err)
	})
	err := c.Visit(url)
	if err != nil {
		return nil, err
	}
	if result.Content == "" {
		return nil, errors.New("未找到回答内容")
	}

	return zhihuTostring(&result), nil
}

// 标准返回格式 title content img[0]
func zhihuTostring(z *zhihu) (res []string) {
	res = append(res, z.Title, z.Content, z.Img[0])
	return
}
