package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/zakislaboratory/rss-to-mdx/pkg/markdown"
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

		md := markdown.NewDocument(item.Content)

		// Remove beehiiv branding
		beehiiv := regexp.MustCompile(`\[Powered by beehiiv\]\(([^)]+)\)`)

		md.RemoveMatches(beehiiv)

		content, err := md.Content()
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

		// Add any additional content to the end of the MDX file
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
