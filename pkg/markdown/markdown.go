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

	var content string

	doc.Find("body").Children().Each(func(_ int, s *goquery.Selection) {

		// Convert the selection to Markdown
		markdownContent := convertSelection(s)

		// If the selection should be removed, skip it
		if md.shouldRemove(markdownContent) {
			return
		}

		// Add the Markdown content to the document
		content += markdownContent + "\n\n"
	})

	content = strings.ReplaceAll(content, "\n\n", "\n")
	content = strings.TrimSuffix(content, "\n")

	return content, nil
}

func convertSelection(s *goquery.Selection) string {

	elements := make([]Element, 0)

	text := s.Text()

	switch s.Get(0).Data {
	case "p":
		elements = append(elements, NewParagraph(s))
	case "h1", "h2", "h3", "h4", "h5", "h6":
		elements = append(elements, NewHeading(s))
	case "a":
		elements = append(elements, NewLink(s))
	case "ul", "ol":
		elements = append(elements, NewList(s))
	case "body":
		// Do nothing
	default:
		panic(fmt.Sprintf("Unknown element: %s", s.Get(0).Data))
	}

	// If the selection has children, we have to convert it to a
	// Markdown element, then replace the text in the parent with the
	// child's Markdown text.
	s.Children().Each(func(_ int, s *goquery.Selection) {
		childMarkdown := convertSelection(s)
		// Replace the child's text in the parent's text
		text = strings.ReplaceAll(text, s.Text(), childMarkdown)
	})

	// Replace the selection's text with the Markdown text
	text = strings.ReplaceAll(text, s.Text(), elements[0].Markdown())

	return text
}

func (md *document) shouldRemove(markdownContent string) bool {

	for _, pattern := range md.removePatterns {

		if pattern.MatchString(markdownContent) {
			return true
		}
	}

	return false
}

func (md *document) RemoveMatches(pattern *regexp.Regexp) {
	md.removePatterns = append(md.removePatterns, pattern)
}

func NewMarkdownDocument(rawHtml string) MarkdownDocument {
	return &document{
		html:           rawHtml,
		removePatterns: make([]*regexp.Regexp, 0),
	}
}
