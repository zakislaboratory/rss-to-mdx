package markdown

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Text struct {
	content string
}

func NewText(s *goquery.Selection) *Text {
	return &Text{content: s.Text()}
}

func (t *Text) Trimmed() string {
	return strings.TrimSpace(t.content)
}

func (t *Text) Bold() string {
	bold := "**" + t.Trimmed() + "**"
	return strings.ReplaceAll(t.content, t.Trimmed(), bold)
}

func (t *Text) Italic() string {
	italic := "*" + t.Trimmed() + "*"
	return strings.ReplaceAll(t.content, t.Trimmed(), italic)
}

func (t *Text) Strikethrough() string {
	strikethrough := "~~" + t.Trimmed() + "~~"
	return strings.ReplaceAll(t.content, t.Trimmed(), strikethrough)
}
