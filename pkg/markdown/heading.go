package markdown

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Heading represents a Markdown heading element.
// Example:
// # Heading 1
// ## Heading 2
// ### Heading 3
type Heading struct {
	Level   int
	Content string
}

func (h Heading) Markdown() string {
	return fmt.Sprintf("%s %s", strings.Repeat("#", h.Level), h.Content)
}

func (h Heading) Type() ElementType { return ElementTypeHeading }

func NewHeading(s *goquery.Selection) Element {

	lvl, err := s.Attr("level")
	if !err {
		// Default to level 2
		lvl = "2"
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
