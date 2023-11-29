package site

import (
	store2 "crawlnovel/pkg/down/store"
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var dingdian = SiteA{
	Name:     "顶点小说",
	HomePage: "https://www.booktxt.net/",
	Match: []string{
		`https://www\.booktxt\.net/\d+_\d+/*`,
		`https://www\.booktxt\.net/\d+_\d+/\d+\.html/*`,
	},
	BookInfo: func(body io.Reader) (s *store2.Store, err error) {
		body = transform.NewReader(body, simplifiedchinese.GBK.NewDecoder())
		doc, err := htmlquery.Parse(body)
		if err != nil {
			return
		}

		s = &store2.Store{}

		nodeTitle := htmlquery.FindOne(doc, `//div[@id="info"]/h1`)
		s.BookName = htmlquery.InnerText(nodeTitle)

		nodeDesc := htmlquery.Find(doc, `//*[@id="intro"]/p`)
		if len(nodeDesc) == 0 {
			err = fmt.Errorf("no matching desc")
			return
		}
		s.Description = strings.Replace(
			htmlquery.OutputHTML(nodeDesc[0], false),
			"<br/>", "\n",
			-1)

		var author = htmlquery.Find(doc, `//*[@id="info"]/p[1]`)
		s.Author = strings.TrimLeft(htmlquery.OutputHTML(author[0], false), "作\u00a0\u00a0\u00a0\u00a0者：")

		nodeContent := htmlquery.Find(doc, `//*[@id="list"]/dl/dd/a`)
		if len(nodeDesc) == 0 {
			err = fmt.Errorf("no matching contents")
			return
		}

		var vol = store2.Volume{
			Name:     "正文",
			Chapters: make([]store2.Chapter, 0),
		}

		for _, v := range nodeContent[9:] {
			//fmt.Printf("href: %v\n", chapter_u)
			chapterURL, err := url.Parse(htmlquery.SelectAttr(v, "href"))
			if err != nil {
				return nil, err
			}

			vol.Chapters = append(vol.Chapters, store2.Chapter{
				Name: strings.TrimSpace(htmlquery.InnerText(v)),
				URL:  chapterURL.String(),
			})
		}
		s.Volumes = append(s.Volumes, vol)
		s.CoverURL = htmlquery.SelectAttr(htmlquery.FindOne(doc, `//*[@id="fmimg"]/img`), "src")
		return
	},
	Chapter: func(body io.Reader) ([]string, error) {
		body = transform.NewReader(body, simplifiedchinese.GBK.NewDecoder())
		doc, err := htmlquery.Parse(body)
		if err != nil {
			return nil, err
		}

		var M []string
		//list
		nodeContent := htmlquery.Find(doc, `//*[@id="content"]/p`)
		nodeContent = append(nodeContent, htmlquery.Find(doc, `//*[@id="content"]/text()`)...)
		if len(nodeContent) == 0 {
			err = fmt.Errorf("No matching content")
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
