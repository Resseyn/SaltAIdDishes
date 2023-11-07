package scrappers

import (
	"SaltAIdDishes/pkg/loggers"
	"fmt"
	"github.com/gocolly/colly/v2"
	"net/url"
)

func Scrap(name string) string {

	encodedName := url.QueryEscape(name)
	c := colly.NewCollector()
	urls := make([]string, 1, 2)
	url := fmt.Sprintf("https://www.freepik.com/search?format=search&last_filter=query&last_value=%s&query=%s&type=photo", encodedName, encodedName)
	//prev := "div.list-content > section.showcase.showcase--completed > figure.showcase__item.js-detail-data.caption.showcase__item--buttons(1) > div:showcase__content.tags-links > a.showcase__link.js-detail-data-link > img[src]"
	c.OnHTML("img.landscape", func(e *colly.HTMLElement) {
		imageURL := e.Attr("data-src")
		if urls[0] == "" {
			urls[0] = imageURL
		}
	})
	c.OnError(func(r *colly.Response, e error) {
		loggers.ErrorLogger.Println(r.Request.URL, string(r.Body))
		fmt.Println("error:", e, r.Request.URL, string(r.Body))
	})

	err := c.Visit(url)
	if err != nil {
		panic(err)
	}
	return urls[0]
}
