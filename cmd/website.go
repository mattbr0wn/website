package main

import (
	"fmt"
	"log"

	"github.com/mattbr0wn/website/config"
	"github.com/mattbr0wn/website/internal/markdown"
	"github.com/mattbr0wn/website/internal/rss"
	"github.com/mattbr0wn/website/internal/ssg"
)

func main() {
	fmt.Println("Setting up build...")
	ssg.SetupStaticPageBuild()

	fmt.Println("Collecting markdown files...")
	markdownFiles, err := markdown.GetMarkdownFilePaths(config.CONTENT_DIR)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Building static pages...")
	ssg.BuildStaticPages(markdownFiles)

	fmt.Println("Building RSS feed...")
	rss.WriteRssFeed(markdownFiles)

	fmt.Println("Build complete.")
}
