package markdown

import "github.com/PuerkitoBio/goquery"

type Link struct {
	Text string
	Href string
}

func (l *Link) Markdown() string {
	return "[" + l.Text + "](" + l.Href + ")"
}

// convertAnchor converts an HTML anchor element to a Markdown link.
func convertAnchor(s *goquery.Selection) MarkdownElement {
	href, _ := s.Attr("href")
	return &Link{
		Text: s.Text(),
		Href: href,
	}
}
