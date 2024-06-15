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

// Tests todo
func SetupStaticPageBuild() {
	fmt.Println("Setting up build...")
	// Remove the existing "static" directory
	if err := deleteDirectory(config.ROOT_DIR); err != nil {
		log.Fatalf("Error deleting %s: %v", config.ROOT_DIR, err)
	}

	// Create the "static" and "static/img" directories
	static_img_dir := filepath.Join(config.ROOT_DIR, "img")
	if err := createDirectory(static_img_dir); err != nil {
		log.Fatalf("Error creating %s: %v", static_img_dir, err)
	}

	// Copy image files into static
	if err := copyDirectoryContents(config.IMG_DIR, config.ROOT_DIR); err != nil {
		log.Fatalf("Error copying %s into %s: %v", config.IMG_DIR, config.ROOT_DIR, err)
	}
}

// Tests todo
func BuildStaticPages(markdownFiles []string) {
	createStaticDirs(markdownFiles)
	generate404()

	articleData := []markdown.ArticleData{}

	for _, file := range markdownFiles {
		switch file {
		case filepath.Join(config.CONTENT_DIR, config.INDEX):
			generateHtmlPage("index", file, &articleData)
		case filepath.Join(config.CONTENT_DIR, "about", config.INDEX):
			generateHtmlPage("about", file, &articleData)
		case filepath.Join(config.CONTENT_DIR, "writing", config.INDEX):
			// do nothing
		default:
			generateHtmlPage("writing", file, &articleData)
		}
	}
	generateHtmlPage("writing-index", filepath.Join(config.CONTENT_DIR, "writing", config.INDEX), &articleData)
}

func GenerateStaticPath(filePath string) (string, error) {
	if filepath.Ext(filePath) != ".md" {
		return "", fmt.Errorf("Invalid file extension: %s.", filepath.Ext(filePath))
	}

	if filepath.Base(filePath) == config.INDEX {
		filePath = filepath.Join(filepath.Dir(filePath), "index.md")
	}
	trimmedPath := strings.TrimPrefix(filePath, config.CONTENT_DIR)
	trimmedPath = strings.TrimSuffix(trimmedPath, filepath.Ext(trimmedPath)) + ".html"
	staticUrl := filepath.Join(config.ROOT_DIR, trimmedPath)
	return staticUrl, nil
}

// create directory structure for the static site
// Tests todo
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

// Tests todo
func generate404() {
	staticUrl := filepath.Join(config.ROOT_DIR, "404.html")

	f, createErr := createFile(staticUrl)
	if createErr != nil {
		log.Fatalf("Error creating 404.html file: %v", createErr)
	}

	genErr := components.NotFound().Render(context.Background(), f)
	if genErr != nil {
		log.Fatalf("Error generating html for 404 page: %v", genErr)
	}
}

// Tests todo
func generateHtmlPage(contentType string, filePath string, articleData *[]markdown.ArticleData) {
	staticUrl, pathErr := GenerateStaticPath(filePath)
	if pathErr != nil {
		log.Println(pathErr)
	}

	f, createErr := createFile(staticUrl)
	if createErr != nil {
		log.Fatalf("Error creating %s: %v", staticUrl, createErr)
	}

	parsedMarkdown, parseErr := markdown.ParseMarkdownFile(filePath)
	if parseErr != nil {
		log.Fatal(parseErr)
	}

	switch contentType {
	case "index":
		err := components.Index(parsedMarkdown.Html).Render(context.Background(), f)
		if err != nil {
			log.Fatalf("Error generating index.html for %s: %v", staticUrl, err)
		}

	case "about":
		err := components.About(parsedMarkdown.Frontmatter, parsedMarkdown.Html).Render(context.Background(), f)
		if err != nil {
			log.Fatalf("Error generating html for about: %v", err)
		}

	case "writing":
		article := markdown.ArticleData{
			Metadata: parsedMarkdown.Frontmatter,
			Body:     parsedMarkdown.Content,
			Path:     strings.TrimPrefix(staticUrl, config.ROOT_DIR),
		}

		*articleData = append(*articleData, article)
		err := components.Article(parsedMarkdown).Render(context.Background(), f)
		if err != nil {
			log.Fatalf("Error generating html for %s: %v", staticUrl, err)
		}

	case "writing-index":
		err := components.Writing(articleData).Render(context.Background(), f)
		if err != nil {
			log.Fatalf("Error generating html for writing/index.html: %v", err)
		}

	default:
		log.Fatalf("Page type %s not supported", contentType)
	}
}
