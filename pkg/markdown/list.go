package markdown

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func convertList(s *goquery.Selection) string {
	switch s.Get(0).Data {
	case "ul":
		return convertUnorderedList(s)
	case "ol":
		return convertOrderedList(s)
	default:
		panic(fmt.Sprintf("Unknown list type: %s", s.Get(0).Data))
	}
}

func convertUnorderedList(s *goquery.Selection) string {
	var markdownContent string

	s.Find("li").Each(func(_ int, li *goquery.Selection) {
		markdownContent += fmt.Sprintf("* %s\n", li.Text())
	})

	return markdownContent
}

func convertOrderedList(s *goquery.Selection) string {
	var markdownContent string

	s.Find("li").Each(func(i int, li *goquery.Selection) {
		markdownContent += fmt.Sprintf("%d. %s\n", i+1, li.Text())
	})

	return markdownContent
}
