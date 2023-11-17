package markdown

import "github.com/PuerkitoBio/goquery"

type Paragraph struct {
	Content string
}

func (p Paragraph) Markdown() string { return p.Content }

func (p Paragraph) Type() ElementType { return ElementTypeParagraph }

func NewParagraph(s *goquery.Selection) Element {
	return &Paragraph{
		Content: s.Text(),
	}
}
