package ssg

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/mattbr0wn/website/config"
	"github.com/mattbr0wn/website/internal/components"
	"github.com/mattbr0wn/website/internal/markdown"
)

func SetupStaticPageBuild() error {
	fmt.Println("Setting up build...")
	// Remove the existing "static" directory
	if err := deleteDirectory(config.ROOT_DIR); err != nil {
		return err
	}

	// Create the "static" and "static/img" directories
	static_img_dir := filepath.Join(config.ROOT_DIR, "img")
	if err := createDirectory(static_img_dir); err != nil {
		return err
	}

	// Copy image files into static
	if err := copyDirectoryContents(config.IMG_DIR, config.ROOT_DIR); err != nil {
		return err
	}

	return nil
}

func BuildStaticPages(markdownFiles []string) {
	createStaticDirs(markdownFiles)
	create404()

	articleData := []markdown.ArticleData{}

	for _, file := range markdownFiles {
		switch file {
		case filepath.Join(config.CONTENT_DIR, "_index.md"):
			generateHtmlPage("index", file, &articleData)
		case filepath.Join(config.CONTENT_DIR, "about/_index.md"):
			generateHtmlPage("about", file, &articleData)
		case filepath.Join(config.CONTENT_DIR, "writing/_index.md"):
			// do nothing
		default:
			generateHtmlPage("writing", file, &articleData)
		}
	}
	generateHtmlPage("writing-index", filepath.Join(config.CONTENT_DIR, "writing/_index.md"), &articleData)
}

func GenerateStaticUrl(filePath string) string {
	if filepath.Base(filePath) == "_index.md" {
		filePath = filepath.Join(filepath.Dir(filePath), "index.md")
	}
	trimmedPath := strings.TrimPrefix(filePath, config.CONTENT_DIR)
	trimmedPath = strings.TrimSuffix(trimmedPath, filepath.Ext(trimmedPath)) + ".html"
	staticUrl := filepath.Join(config.ROOT_DIR, trimmedPath)
	return staticUrl
}

// create directory structure for the static site
func createStaticDirs(contentFiles []string) error {
	for _, path := range contentFiles {
		// Remove the content dir prefix to get the content path without project structure
		trimmedPath := strings.TrimPrefix(path, config.CONTENT_DIR)

		// Get the directory path by removing the file name
		dirPath := filepath.Dir(trimmedPath)

		// Create the directory if it doesn't exist
		if err := createDirectory(filepath.Join(config.ROOT_DIR, dirPath)); err != nil {
			return err
		}
	}
	return nil
}

func create404() {
	staticUrl := filepath.Join(config.ROOT_DIR, "404.html")
	f, _ := createFile(staticUrl)
	err := components.NotFound().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("ERROR: Failed to create 404 page: %v", err)
	}
}

func generateHtmlPage(contentType string, filePath string, articleData *[]markdown.ArticleData) {
	staticUrl := GenerateStaticUrl(filePath)
	f, _ := createFile(staticUrl)

	metadata, body, mdString := markdown.ParseMarkdownFile(filePath)

	switch contentType {
	case "index":
		err := components.Index(body).Render(context.Background(), f)
		if err != nil {
			log.Fatalf("ERROR: Failed to write index page: %v", err)
		}

	case "about":
		err := components.About(metadata, body).Render(context.Background(), f)
		if err != nil {
			log.Fatalf("ERROR: Failed to write about page: %v", err)
		}

	case "writing":
		article := markdown.ArticleData{
			Metadata: metadata,
			Body:     mdString,
			Path:     strings.TrimPrefix(staticUrl, config.ROOT_DIR),
		}

		*articleData = append(*articleData, article)
		err := components.Article(metadata, body, mdString).Render(context.Background(), f)
		if err != nil {
			log.Fatalf("ERROR: Failed to write %s: %v", staticUrl, err)
		}

	case "writing-index":
		err := components.Writing(articleData).Render(context.Background(), f)
		if err != nil {
			log.Fatalf("ERROR: Failed to write writing index page: %v", err)
		}

	default:
		log.Fatalf("ERROR: Page type not supported")
	}
}
