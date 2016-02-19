package parser

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
)

func XPath() (interface{}, error) {
	doc, err := goquery.NewDocument("http://www.guokr.com/")
	if err != nil {
		return nil, err
	}
	doc.Find(".main a").Each(func(i int, s *goquery.Selection) {
		title, _ := s.Attr("title")
		alt, _ := s.Find("img").Attr("alt")
		fmt.Printf("%s %s\n", title, alt)
	})
	return nil, nil
}
