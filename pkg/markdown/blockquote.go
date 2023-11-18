package markdown

import "github.com/PuerkitoBio/goquery"

type BlockQuote struct {
	content string
}

func NewBlockQuote(s *goquery.Selection) *BlockQuote {
	return &BlockQuote{content: s.Text()}
}

func (bq *BlockQuote) Markdown() string {
	return "> " + bq.content
}
