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

func SetupStaticPageBuild() {
	if err := deleteDirectory(config.ROOT_DIR); err != nil {
		log.Fatalf("Error deleting %s: %v", config.ROOT_DIR, err)
	}

	staticImgDir := filepath.Join(config.ROOT_DIR, "img")
	if err := createDirectories(config.ROOT_DIR, staticImgDir); err != nil {
		log.Fatalf("Error creating img directory: %v", err)
	}

	if err := copyImages(); err != nil {
		log.Fatalf("Error copying img directory: %v", err)
	}
}

func BuildStaticPages(markdownFiles []string) {
	if err := createStaticDirs(markdownFiles); err != nil {
		log.Fatalf("Error creating static directories: %v", err)
	}

	if err := generate404(); err != nil {
		log.Fatalf("Error creating 404.html: %v", err)
	}

	articleData := []markdown.ParsedMarkdownFile{}

	for _, file := range markdownFiles {
		switch file {
		case filepath.Join(config.CONTENT_DIR, config.INDEX):
			if err := generateIndexPage(file, &articleData); err != nil {
				log.Fatalf("Error generating index page: %v", err)
			}
		case filepath.Join(config.CONTENT_DIR, "about", config.INDEX):
			if err := generateAboutPage(file, &articleData); err != nil {
				log.Fatalf("Error generating about page: %v", err)
			}
		case filepath.Join(config.CONTENT_DIR, "writing", config.INDEX):
			// do nothing
		default:
			if err := generateWritingPage(file, &articleData); err != nil {
				log.Fatalf("Error generating writing page: %v", err)
			}
		}
	}

	if err := generateWritingIndexPage(filepath.Join(config.CONTENT_DIR, "writing", config.INDEX), &articleData); err != nil {
		log.Fatalf("Error generating writing/index.html: %v", err)
	}
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

func createDirectories(dirs ...string) error {
	for _, dir := range dirs {
		if err := createDirectory(dir); err != nil {
			return fmt.Errorf("Error creating directory %s: %v", dir, err)
		}
	}
	return nil
}

func copyImages() error {
	if err := copyDirectoryContents(config.IMG_DIR, config.ROOT_DIR); err != nil {
		return fmt.Errorf("Error copying %s into %s: %v", config.IMG_DIR, config.ROOT_DIR, err)
	}
	return nil
}

func createStaticDirs(contentFiles []string) error {
	for _, path := range contentFiles {
		trimmedPath := strings.TrimPrefix(path, config.CONTENT_DIR)
		dirPath := filepath.Dir(trimmedPath)
		if err := createDirectory(filepath.Join(config.ROOT_DIR, dirPath)); err != nil {
			return err
		}
	}
	return nil
}

func generate404() error {
	staticUrl := filepath.Join(config.ROOT_DIR, "404.html")

	f, err := createFile(staticUrl)
	if err != nil {
		return fmt.Errorf("Error creating 404.html file: %v", err)
	}

	if err := components.NotFound().Render(context.Background(), f); err != nil {
		return fmt.Errorf("Error generating html for 404 page: %v", err)
	}

	return nil
}

func generateIndexPage(filePath string, articleData *[]markdown.ParsedMarkdownFile) error {
	return generateHtmlPage("index", filePath, articleData)
}

func generateAboutPage(filePath string, articleData *[]markdown.ParsedMarkdownFile) error {
	return generateHtmlPage("about", filePath, articleData)
}

func generateWritingPage(filePath string, articleData *[]markdown.ParsedMarkdownFile) error {
	return generateHtmlPage("writing", filePath, articleData)
}

func generateWritingIndexPage(filePath string, articleData *[]markdown.ParsedMarkdownFile) error {
	return generateHtmlPage("writing-index", filePath, articleData)
}

func generateHtmlPage(contentType string, filePath string, articleData *[]markdown.ParsedMarkdownFile) error {
	staticUrl, err := GenerateStaticPath(filePath)
	if err != nil {
		return err
	}

	f, err := createFile(staticUrl)
	if err != nil {
		return fmt.Errorf("Error creating %s: %v", staticUrl, err)
	}

	parsedMarkdown, err := markdown.ParseMarkdownFile(filePath)
	if err != nil {
		return err
	}

	parsedMarkdown.SetStaticFileUrl(strings.TrimPrefix(staticUrl, config.ROOT_DIR))

	switch contentType {
	case "index":
		html := components.Unsafe(parsedMarkdown.BodyAsString())
		if err := components.Index(html).Render(context.Background(), f); err != nil {
			return fmt.Errorf("Error generating index.html for %s: %v", staticUrl, err)
		}

	case "about":
		html := components.Unsafe(parsedMarkdown.BodyAsString())
		if err := components.About(parsedMarkdown.Frontmatter(), html).Render(context.Background(), f); err != nil {
			return fmt.Errorf("Error generating html for about: %v", err)
		}

	case "writing":
		*articleData = append(*articleData, *parsedMarkdown)
		if err := components.Article(*parsedMarkdown).Render(context.Background(), f); err != nil {
			return fmt.Errorf("Error generating html for %s: %v", staticUrl, err)
		}

	case "writing-index":
		if err := components.Writing(articleData).Render(context.Background(), f); err != nil {
			return fmt.Errorf("Error generating html for writing/index.html: %v", err)
		}

	default:
		return fmt.Errorf("Page type %s not supported", contentType)
	}

	return nil
}
