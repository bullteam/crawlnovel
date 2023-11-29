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

var biquge3 = SiteA{
	Name:     "笔趣阁",
	HomePage: "https://www.biqivge.com/",
	Match: []string{
		`https://www\.biqivge\.com/book/\d+/*`,
		`https*://www\.biqivge\.com/book/goto/id/\d+/*`,
		`https://www\.biqivge\.com/book/\d+/\d+\.html/*`,
	},
	BookInfo: func(body io.Reader) (s *store2.Store, err error) {
		body = transform.NewReader(body, simplifiedchinese.GBK.NewDecoder())
		doc, err := htmlquery.Parse(body)
		if err != nil {
			return
		}

		s = &store2.Store{}

		s.BookName = htmlquery.InnerText(htmlquery.FindOne(doc, `//div[@class="info"]/h2`))

		// var author = htmlquery.Find(doc, `//*[@id="info"]/p[1]`)
		raw_author := htmlquery.InnerText(htmlquery.FindOne(doc, `//div[@class="small"]/span[1]`))
		s.Author = strings.TrimSpace(strings.TrimLeft(raw_author, "作者："))

		node_content := htmlquery.Find(doc, `//div[@class="listmain"]/dl/dd/a`)
		if len(node_content) == 0 {
			err = fmt.Errorf("No matching contents")
			return
		}

		var vol = store2.Volume{
			Name:     "正文",
			Chapters: make([]store2.Chapter, 0),
		}

		for _, v := range node_content[6:] {
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

		return
	},
	Chapter: func(body io.Reader) ([]string, error) {
		body = transform.NewReader(body, simplifiedchinese.GBK.NewDecoder())
		doc, err := htmlquery.Parse(body)
		if err != nil {
			return nil, err
		}

		M := []string{}
		//list
		nodeContent := htmlquery.Find(doc, `//*[@id="content"]/text()`)
		if len(nodeContent) == 0 {
			err = fmt.Errorf("no matching content")
			return nil, err
		}
		for _, v := range nodeContent {
			t := htmlquery.InnerText(v)
			t = strings.TrimSpace(t)

			if strings.HasPrefix(t, "…") {
				continue
			}

			t = strings.Replace(t, "…", "", -1)
			t = strings.Replace(t, ".......", "", -1)
			t = strings.Replace(t, "...”", "”", -1)

			if t == "" {
				continue
			}

			M = append(M, t)
		}
		return M, nil
	},
	Search: func(s string) (result []ChaperSearchResult, err error) {
		baseurl, err := url.Parse("https://so.biqusoso.com/s.php")
		if err != nil {
			return
		}
		value := baseurl.Query()
		value.Add("ie", "utf-8")
		value.Add("siteid", "biqiuge.com")
		value.Add("s", "2758772450457967865")
		value.Add("q", s)
		baseurl.RawQuery = value.Encode()

		resp, err := RequestGet(baseurl.String())
		if err != nil {
			return
		}
		defer resp.Body.Close()
		body := resp.Body
		// body := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
		doc, err := htmlquery.Parse(body)
		if err != nil {
			return
		}

		r := htmlquery.Find(doc, `//*[@class="search-list"]/ul/li`)
		if len(r) == 0 {
			return nil, nil
		}
		for _, v := range r[1:] {
			s2 := htmlquery.FindOne(v, `/span[2]/a`)
			s4 := htmlquery.FindOne(v, `/span[3]`)
			r := ChaperSearchResult{
				BookName: htmlquery.InnerText(s2),
				Author:   htmlquery.InnerText(s4),
				BookURL:  htmlquery.SelectAttr(s2, "href"),
			}
			result = append(result, r)
		}
		return
	},
}
