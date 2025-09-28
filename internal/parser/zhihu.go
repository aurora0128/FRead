package parser

import (
	"errors"
	"log"
	"ppeua/FRead/internal/config"
	"strings"
	"time"

	htmlToMarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

type zhihu struct {
	Url   string `json:"url"`
	Title string `json:"question"`
	//content为markdown格式
	Content string   `json:"content"`
	Img     []string `json:"img"`
}

func ParserUrlZhihu(url string, markdownPath string) ([]string, error) {
	c := colly.NewCollector()
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*zhihu.com",
		Parallelism: 1,
		Delay:       2 * time.Second,
	})
	var result zhihu
	result.Img = make([]string, 0)
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
		htmlContent, err := e.DOM.Html()
		if err != nil {
			log.Printf("HTML解析失败: %v\n", err)
			return
		}

		// 使用goquery解析HTML
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
		if err != nil {
			log.Printf("goquery解析失败: %v\n", err)
			return
		}

		// 找到所有图片并替换src属性
		doc.Find("img").Each(func(i int, s *goquery.Selection) {
			// 检查是否有data-original属性
			if originalSrc, exists := s.Attr("data-original"); exists && originalSrc != "" {
				// 直接设置src属性为真实图片URL
				s.SetAttr("src", originalSrc)
				//将图片url添加到result.Img
				result.Img = append(result.Img, originalSrc)
				log.Printf("替换图片 %d: %s", i+1, originalSrc)
			}
		})

		// 获取修改后的HTML
		modifiedHTML, err := doc.Html()
		if err != nil {
			log.Printf("获取修改后HTML失败: %v\n", err)
			return
		}

		makrdown, err := htmlToMarkdown.ConvertString(modifiedHTML)
		if err != nil {
			log.Printf("Markdown解析失败: %v\n", err)
			return
		}

		result.Content = makrdown
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
