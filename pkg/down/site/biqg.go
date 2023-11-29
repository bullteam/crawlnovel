package site

import (
	store2 "crawlnovel/pkg/down/store"
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/antchfx/htmlquery"
)

var bigq = SiteA{
	Name:     "笔趣阁",
	HomePage: "https://www.biqg.cc/",
	Match: []string{
		`https://www\.biqg\.cc/book/\d+/*`,
		`https://www\.biqg\.cc/book/\d+/\d+\.html/*`,
	},
	BookInfo: func(body io.Reader) (s *store2.Store, err error) {
		//body = transform.NewReader(body, simplifiedchinese.GBK.NewDecoder())
		doc, err := htmlquery.Parse(body)
		if err != nil {
			return
		}
		s = &store2.Store{}

		s.BookName = htmlquery.InnerText(htmlquery.FindOne(doc, `//div[@class="info"]/h1`))

		rawAuthor := htmlquery.InnerText(htmlquery.FindOne(doc, `//div[@class="small"]/span[1]`))
		s.Author = strings.TrimSpace(strings.TrimLeft(rawAuthor, "作者："))

		nodeContent := htmlquery.Find(doc, `//div[@class="listmain"]/dl/dd/a`)

		if len(nodeContent) == 0 {
			err = fmt.Errorf("no matching contents")
			return
		}

		var vol = store2.Volume{
			Name:     "正文",
			Chapters: make([]store2.Chapter, 0),
		}

		for _, v := range nodeContent {
			chapterURL, err := url.Parse(htmlquery.SelectAttr(v, "href"))
			fmt.Printf("href: %v\n", chapterURL)
			if err != nil {
				return nil, err
			}
			if chapterURL.String() == "javascript:dd_show()" {
				fmt.Println("sssss", chapterURL.String())

				//return nil, nil
			}

			vol.Chapters = append(vol.Chapters, store2.Chapter{
				Name: strings.TrimSpace(htmlquery.InnerText(v)),
				URL:  chapterURL.String(),
			})
		}
		//s.Volumes = append(s.Volumes, vol)

		return
	},
	Chapter: func(body io.Reader) ([]string, error) {
		//body = transform.NewReader(body, simplifiedchinese.GBK.NewDecoder())
		doc, err := htmlquery.Parse(body)
		if err != nil {
			return nil, err
		}

		var M []string
		//list
		nodeContent := htmlquery.Find(doc, `//*[@id="chaptercontent"]/text()`)
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
