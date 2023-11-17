package markdown

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// convertHtml converts HTML to a simplified Markdown format.
func ConvertHtml(htmlContent string) (string, error) {
	// Remove classes and styles from HTML
	htmlContent = stripClassesFromHTML(htmlContent)
	htmlContent = stripStylesFromHTML(htmlContent)

	// Load HTML content into a goquery document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		return "", err
	}
	// Convert HTML elements to Markdown
	var markdownContent string
	doc.Find("body *").Each(func(_ int, s *goquery.Selection) {
		switch s.Get(0).Data {
		case "p":
			markdownContent += "\n" + s.Text() + "\n"
		case "h1", "h2", "h3", "h4", "h5", "h6":
			level, _ := strconv.Atoi(s.AttrOr("level", "2"))
			markdownContent += fmt.Sprintf("\n%s %s\n", strings.Repeat("#", level), s.Text())
		case "a":
			href, _ := s.Attr("href")
			markdownContent += fmt.Sprintf("[%s](%s)", s.Text(), href)
		case "ul":
			markdownContent += "\n"
			s.Find("li").Each(func(_ int, li *goquery.Selection) {
				markdownContent += fmt.Sprintf("* %s\n", li.Text())
			})
			markdownContent += "\n"
		case "ol":
			markdownContent += "\n"
			s.Find("li").Each(func(i int, li *goquery.Selection) {
				markdownContent += fmt.Sprintf("%d. %s\n", i+1, li.Text())
			})
			markdownContent += "\n"
		default:
			// markdownContent += "\n" + s.Text() + "\n"
		}
	})

	// Define the regex pattern for the "Powered by beehiiv" line
	poweredByBeehiivPattern := `\[[^\]]+\]\(https:\/\/www\.beehiiv\.com\/\?.*?\)`

	// Compile the regex pattern
	re := regexp.MustCompile(poweredByBeehiivPattern)

	// Strip "Powered by beehiiv" at the end of the content
	markdownContent = re.ReplaceAllString(markdownContent, "")

	return markdownContent, nil
}

// stripClassesFromHTML removes class attributes from HTML elements
// i.e. : <div class='myUnknownClass'> => <div>
func stripClassesFromHTML(htmlContent string) string {
	re := regexp.MustCompile(` class=['"][^'"]*['"]`)
	return re.ReplaceAllString(htmlContent, "")
}

// stripStylesFromHTML removes style attributes from HTML elements
// i.e. : <div style='color: red;'> or <div style="color: red;"> => <div>
func stripStylesFromHTML(htmlContent string) string {
	re := regexp.MustCompile(` style=['"][^'"]*['"]`)
	return re.ReplaceAllString(htmlContent, "")
}
