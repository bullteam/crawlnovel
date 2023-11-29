package site

import (
	"crawlnovel/pkg/down/store"
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"net/url"
	"strings"
)

var dfu7 = SiteA{
	Name:     "夫妻小说网",
	HomePage: "https://www.dfu7.com/",
	Match: []string{
		`https://www\.dfu7\.com/NanRen/Shu_\d+\.Html/*`,
		`https://www\.dfu7\.com/NanRen/\d+\.Html/*`,
		`http://www\.dfu7\.com/NanRen/Chapter_\d+\.Html/*`,
	},
	BookInfo: func(body io.Reader) (s *store.Store, err error) {
		body = transform.NewReader(body, simplifiedchinese.GBK.NewDecoder())
		doc, err := htmlquery.Parse(body)
		if err != nil {
			return
		}
		s = &store.Store{}
		nodeTitle := htmlquery.Find(doc, `//*[@class=="title"]/h1`)
		if len(nodeTitle) == 0 {
			err = fmt.Errorf("no matching title")
			return
		}
		s.BookName = htmlquery.InnerText(nodeTitle[0])

		nodeDesc := htmlquery.Find(doc, `//*[@class="con gray"]/text()`)
		if len(nodeDesc) == 0 {
			err = fmt.Errorf("no matching desc")
			return
		}
		s.Description = strings.Replace(
			htmlquery.OutputHTML(nodeDesc[0], false),
			"<br/>", "\n",
			-1)
		var author = htmlquery.Find(doc, `//*[@class="detail"]/p[1]`)
		s.Author = strings.TrimLeft(htmlquery.OutputHTML(author[0], false), "作\u00a0\u00a0\u00a0\u00a0者：")
		nodeContent := htmlquery.Find(doc, `//*[@class="list"]/dl/dd/a`)
		if len(nodeDesc) == 0 {
			err = fmt.Errorf("no matching contents")
			return
		}
		var vol = store.Volume{
			Name:     "正文",
			Chapters: make([]store.Chapter, 0),
		}
		for _, v := range nodeContent[9:] {
			//fmt.Printf("href: %v\n", chapter_u)
			chapterURL, err := url.Parse(htmlquery.SelectAttr(v, "href"))
			if err != nil {
				return nil, err
			}

			vol.Chapters = append(vol.Chapters, store.Chapter{
				Name: strings.TrimSpace(htmlquery.InnerText(v)),
				URL:  chapterURL.String(),
			})
		}
		s.Volumes = append(s.Volumes, vol)
		s.CoverURL = htmlquery.SelectAttr(htmlquery.FindOne(doc, `//*[@class="info flex"]/img`), "src")
		return
	},
	Chapter: func(body io.Reader) ([]string, error) {
		body = transform.NewReader(body, simplifiedchinese.GBK.NewDecoder())
		doc, err := htmlquery.Parse(body)
		if err != nil {
			return nil, err
		}

		var M []string
		nodeContent := htmlquery.Find(doc, `//div[@id="ak"]/text()`)
		if len(nodeContent) == 0 {
			err = fmt.Errorf("no matching content")
			return nil, err
		}
		for _, v := range nodeContent {
			t := htmlquery.InnerText(v)
			t = strings.TrimSpace(t)
			M = append(M, t)
		}

		return M, nil
	},
}
