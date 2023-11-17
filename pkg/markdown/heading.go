package markdown

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Heading struct {
	Level   int
	Content string
}

func (h Heading) Markdown() string {
	return fmt.Sprintf("%s %s\n", strings.Repeat("#", h.Level), h.Content)
}

func convertHeading(s *goquery.Selection) MarkdownElement {

	lvl, err := s.Attr("level")
	if !err {
		panic(fmt.Sprintf("Heading element does not have a level attribute: %s", s.Text()))
	}

	level := int(lvl[0] - '0')

	if level < 1 || level > 6 {
		panic(fmt.Sprintf("Heading element has an invalid level attribute: %s", s.Text()))
	}

	return &Heading{
		Level:   level,
		Content: s.Text(),
	}
}
