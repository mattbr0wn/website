package ssg

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mattbr0wn/website/config"
	"github.com/mattbr0wn/website/internal/components"
	"github.com/mattbr0wn/website/internal/markdown"
)

func BuildStaticPages(rootPath string, headData config.HeadData) {
	createRootDir(rootPath)

	files, err := GetMarkdownFiles("web/content")
	if err != nil {
		log.Fatalf("ERROR: Unable to get markdown files")
	}

	createContentDirs(rootPath, files)
	create404(rootPath)
	articleData := []markdown.ArticleData{}
	var pageType string

	for _, file := range files {
		switch file {
		case "web/content/index.md":
			pageType = "index"
		case "web/content/about/index.md":
			pageType = "about"
		case "web/content/writing/index.md":
			pageType = ""
		default:
			pageType = "writing"
		}

		if pageType != "" {
			createPage(pageType, rootPath, headData, file, &articleData)
		}
	}
	createPage("writing-index", rootPath, headData, "web/content/writing/index.md", &articleData)
}

func GeneratePath(rootPath string, filePath string) string {
	trimmedPath := strings.TrimPrefix(filePath, "web/content")
	newPath := rootPath + trimmedPath
	newPathWithoutExt := strings.TrimSuffix(newPath, filepath.Ext(newPath))
	path := newPathWithoutExt + ".html"
	return path
}

func GetMarkdownFiles(dir string) ([]string, error) {
	var markdownFiles []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".md" {
			markdownFiles = append(markdownFiles, path)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return markdownFiles, nil
}

func createHtmlFile(path string) *os.File {
	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("ERROR: Failed to create index page: %v", err)
	}
	return f
}

func createRootDir(rootPath string) {
	if _, err := os.Stat(rootPath); os.IsNotExist(err) {
		if err := os.Mkdir(rootPath, 0755); err != nil {
			log.Fatalf("ERROR: Failed to create output directory: %v", err)
		}
	} else if err != nil {
		log.Fatalf("ERROR: Failed to check output directory: %v", err)
	}
}

func createContentDirs(rootPath string, contentFiles []string) error {
	for _, path := range contentFiles {
		// Remove the "content/" prefix from the path
		trimmedPath := strings.TrimPrefix(path, "web/content/")

		// Get the directory path by removing the file name
		dirPath := filepath.Dir(trimmedPath)

		// Create the directory if it doesn't exist
		err := os.MkdirAll(filepath.Join(rootPath, dirPath), os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func create404(rootPath string) {
	name := rootPath + "/404.html"
	f := createHtmlFile(name)
	err := components.NotFound().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("ERROR: Failed to create 404 page: %v", err)
	}
}

func createPage(contentType string, rootPath string, headData config.HeadData, contentPath string, articleData *[]markdown.ArticleData) {
	name := GeneratePath(rootPath, contentPath)
	trimmedName := strings.TrimPrefix(name, "web/static/")

	f := createHtmlFile(name)
	metadata, body, mdString := markdown.ParseMarkdownFile(contentPath)

	switch contentType {
	case "index":
		err := components.Index(headData, body).Render(context.Background(), f)
		if err != nil {
			log.Fatalf("ERROR: Failed to write index page: %v", err)
		}

	case "about":
		err := components.About(headData, metadata, body).Render(context.Background(), f)
		if err != nil {
			log.Fatalf("ERROR: Failed to write about page: %v", err)
		}

	case "writing":
		article := markdown.ArticleData{
			Metadata: metadata,
			Body:     mdString,
			Path:     trimmedName,
		}

		*articleData = append(*articleData, article)
		err := components.Article(headData, metadata, body, mdString).Render(context.Background(), f)
		if err != nil {
			log.Fatalf("ERROR: Failed to write %s: %v", name, err)
		}

	case "writing-index":
		err := components.Writing(headData, articleData).Render(context.Background(), f)
		if err != nil {
			log.Fatalf("ERROR: Failed to write writing index page: %v", err)
		}

	default:
		log.Fatalf("ERROR: Page type not supported")
	}
}
