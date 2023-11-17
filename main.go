package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// RSSItem represents the structure of an RSS item
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

	xmlPath := flag.String("xml", "", "Path to the XML file containing RSS items")
	outputDir := flag.String("outputDir", "output", "Output directory for MDX files")
	flag.Parse()

	if *xmlPath == "" {
		fmt.Println("Error: Please provide the path to the XML file using the -xml flag.")
		return
	}

	xmlData, err := ioutil.ReadFile(*xmlPath)
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
		// Prepare MDX content
		mdxContent := fmt.Sprintf(`---
title: "%s"
description: "%s"
link: "%s"
guid: "%s"
pubDate: "%s"
creator: "%s"
---

%s
`, item.Title, item.Description, item.Link, item.GUID, item.PubDate, item.Creator, item.Content)

		// Clean up title for filename
		filename := cleanFilename(item.Title) + ".mdx"

		// Write MDX content to file
		outputPath := filepath.Join(*outputDir, filename)
		err := ioutil.WriteFile(outputPath, []byte(mdxContent), os.ModePerm)
		if err != nil {
			fmt.Printf("Error writing MDX file %s: %v\n", filename, err)
		} else {
			fmt.Printf("MDX file %s created successfully\n", filename)
		}
	}
}

// cleanFilename removes invalid characters from a string to be used as a filename
func cleanFilename(s string) string {
	// Replace invalid characters with underscores
	invalidChars := []string{"/", ":", "?", "<", ">", "|", "\"", "*"}
	for _, char := range invalidChars {
		s = strings.ReplaceAll(s, char, "_")
	}
	return s
}
