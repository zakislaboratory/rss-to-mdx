package markdown

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type MarkdownDocument interface {
	// Content returns the Markdown representation of the HTML document.
	Content() (string, error)

	// RemoveMatches scans the Markdown document for any elements whose
	// Markdown content matches the given regex pattern and removes them.
	RemoveMatches(*regexp.Regexp)
}

type document struct {
	// The raw HTML content to convert to Markdown.
	html string

	// Patterns to remove from the Markdown document.
	// Useful for removing unwanted content from the Markdown document.
	removePatterns []*regexp.Regexp
}

func (md *document) Content() (string, error) {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(md.html))
	if err != nil {
		return "", fmt.Errorf("error loading HTML content into goquery document: %v", err)
	}

	content := convertSelection(doc.Find("body").Children())

	for _, pattern := range md.removePatterns {
		content = pattern.ReplaceAllString(content, "")
	}

	return content, nil
}

func convertSelection(s *goquery.Selection) string {

	text := s.Text()

	if text == "" {
		return ""
	}

	switch s.Get(0).Data {
	case "p":
		text = NewParagraph(s).Markdown()
	case "h1", "h2", "h3", "h4", "h5", "h6":
		text = NewHeading(s).Markdown()
	case "a":
		text = NewLink(s).Markdown()
	case "ul", "ol":
		text = NewList(s).Markdown()
	case "li":
		text = NewListItem(s).Markdown()
	case "strong", "b":
		text = NewText(s).Bold()
	case "em", "i":
		text = NewText(s).Italic()
	case "del":
		text = NewText(s).Strikethrough()
	case "blockquote":
		text = NewBlockQuote(s).Markdown()
	case "code":
		text = NewCodeBlock(s).Markdown()
	default:
		// Do nothing
	}

	// If the selection has children, we have to convert it to a
	// Markdown element, then replace the text in the parent with the
	// child's Markdown text.
	s.Children().Each(func(_ int, c *goquery.Selection) {

		childMarkdown := convertSelection(c)

		// Replace the child's text in the parent's text
		if childMarkdown != "" {
			text = strings.Replace(text, c.Text(), childMarkdown, 1)
		}
	})

	// We only add space for certain elements
	el := s.Get(0).Data

	if el == "p" || el == "li" {
		text = text + "\n\n"
	}

	// for blockquotes and headings, we also prepend a newline
	if el == "blockquote" || el == "h1" || el == "h2" || el == "h3" || el == "h4" || el == "h5" || el == "h6" {
		text = "\n\n" + text + "\n\n"
	}

	return text
}

func (md *document) RemoveMatches(pattern *regexp.Regexp) {
	md.removePatterns = append(md.removePatterns, pattern)
}

func NewDocument(rawHtml string) MarkdownDocument {
	return &document{
		html:           rawHtml,
		removePatterns: make([]*regexp.Regexp, 0),
	}
}
