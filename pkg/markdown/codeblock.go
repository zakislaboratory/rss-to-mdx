package markdown

import "github.com/PuerkitoBio/goquery"

type CodeBlock struct {
	content string
}

func NewCodeBlock(s *goquery.Selection) *CodeBlock {
	return &CodeBlock{content: s.Text()}
}

func (cb *CodeBlock) Markdown() string {
	return "```\n" + cb.content + "\n```"
}
