package markdown

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type List struct {
	Content   string
	IsOrdered bool
}

func (l *List) Markdown() string { return l.Content }

func (l *List) Type() ElementType { return ElementTypeList }

func NewList(s *goquery.Selection) Element {

	switch s.Get(0).Data {

	case "ul":
		return &List{convertUnorderedList(s), false}
	case "ol":
		return &List{convertOrderedList(s), true}

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

type ListItem struct {
	Content string
}

func (li *ListItem) Markdown() string { return li.Content }

func (li *ListItem) Type() ElementType { return ElementTypeListItem }

func NewListItem(s *goquery.Selection) Element {
	return &ListItem{convertListItem(s)}
}

func convertListItem(s *goquery.Selection) string {
	return s.Text()
}
