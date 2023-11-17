package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type RSSItem struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	GUID        string `xml:"guid"`
	PubDate     string `xml:"pubDate"`
	Creator     string `xml:"creator"`
	Content     string `xml:"encoded"`
}

func main() {

	showHelp := flag.Bool("help", false, "Show help and usage information")
	xmlPath := flag.String("xml", "", "Path to the XML file containing RSS items")
	outputDir := flag.String("out", "output", "Output directory for MDX files")
	flag.Parse()

	if *showHelp {
		flag.Usage()
		return
	}

	if *xmlPath == "" {
		fmt.Println("Error: Please provide the path to the XML file using the -xml flag.")
		return
	}

	xmlData, err := os.ReadFile(*xmlPath)
	if err != nil {
		fmt.Printf("Error reading XML file: %v\n", err)
		return
	}

	var rss struct {
		Items []RSSItem `xml:"channel>item"`
	}
	err = xml.Unmarshal(xmlData, &rss)
	if err != nil {
		fmt.Printf("Error unmarshalling XML: %v\n", err)
		return
	}

	err = os.MkdirAll(*outputDir, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		return
	}

	// Process each RSS item
	for _, item := range rss.Items {

		content := stripClassesFromHTML(item.Content)
		content = stripStylesFromHTML(content)
		content, err = convertHTMLToMarkdownSimple(content)
		if err != nil {
			fmt.Printf("Error converting HTML to Markdown: %v\n", err)
			return
		}

		date, err := convertDateFormat(item.PubDate)
		if err != nil {
			fmt.Printf("Error converting date format: %v\n", err)
			return
		}

		mdxContent := fmt.Sprintf(`---
title: "%s"
description: "%s"
link: "%s"
guid: "%s"
date: "%s"
lang: "en"
tags: ["Personal Development", "Systems"]
category: "Growth"
creator: "%s"
---

%s
`, item.Title, item.Description, item.Link, item.GUID,
			date, item.Creator, content)

		// Add a line at the bottom that says "this was first published on Circadian Growth"
		mdxContent += ("\n\n---\n\nThis was originally published on my weekly newsletter [Circadian Growth](circadiangrowth.com)")

		// Clean up title for filename
		filename := cleanFilename(item.Title) + ".mdx"

		// Write MDX content to file
		outputPath := filepath.Join(*outputDir, filename)
		err = os.WriteFile(outputPath, []byte(mdxContent), os.ModePerm)
		if err != nil {
			fmt.Printf("Error writing MDX file %s: %v\n", filename, err)
		} else {
			fmt.Printf("MDX file %s created successfully\n", filename)
		}
	}
}

// convertHTMLToMarkdownSimple converts HTML to a simplified Markdown format.
func convertHTMLToMarkdownSimple(htmlContent string) (string, error) {
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

// cleanFilename removes invalid characters from a string to be used as a filename
func cleanFilename(s string) string {

	invalidChars := []string{"/", ":", "?", "<", ">", "|", "\"", "*", "'", ",", "&"}

	for _, char := range invalidChars {
		s = strings.ReplaceAll(s, char, "")
	}

	s = strings.ReplaceAll(s, " ", "-")
	s = strings.Trim(s, "_")
	s = strings.ToLower(s)

	return s
}

func convertDateFormat(inputDate string) (string, error) {
	parsedTime, err := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", inputDate)
	if err != nil {
		return "", err
	}

	outputDate := parsedTime.Format("2006-01-02")

	return outputDate, nil
}
