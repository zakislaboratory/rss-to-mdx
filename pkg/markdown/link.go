package markdown

import "github.com/PuerkitoBio/goquery"

type Link struct {
	Text string
	Href string
}

func (l *Link) Markdown() string {
	return "[" + l.Text + "](" + l.Href + ")"
}

func (l *Link) Type() ElementType { return ElementTypeAnchor }

// NewLink creates a Markdown link from a *goquery.Selection.
func NewLink(s *goquery.Selection) Element {
	href, _ := s.Attr("href")
	return &Link{
		Text: s.Text(),
		Href: href,
	}
}
