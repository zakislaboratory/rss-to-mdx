# RSS to MDX Converter

Convert RSS feed items to individual MDX (Markdown with Front Matter) files for static site generators or MDX-based bloggs.

## How it Works

- Parses an XML file containing RSS feed items.
- Outputs each RSS item as a separate MDX file in a specified directory.
- Generates front matter with metadata such as title, description, link, GUID, pubDate, and creator.
- Cleans up titles for use as filenames.


## Installation


Clone this repository:

```bash
git clone https://github.com/your-username/rss-to-mdx-converter.git
cd rss-to-mdx-converter
```
  

## Usage 

```
go run main.go -xml path/to/rss.xml -outputDir path/to/output
```
  
- xml: Path to the XML file containing RSS feed items.
- outputDir (optional): Output directory for MDX files. Default is output in the current directory.



