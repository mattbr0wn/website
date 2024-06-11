package markdown

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/a-h/templ"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/frontmatter"
)

type ArticleData struct {
	Metadata Frontmatter
	Body     *string
	Path     string
}

type Frontmatter struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Date        string `yaml:"date"`
	HeroImg     string `yaml:"hero"`
	Draft       bool   `yaml:"draft"`
}

type ParsedMarkdownFile struct {
	Frontmatter Frontmatter
	Html        templ.Component
	Content     *string
}

// Returns an array of filepaths for markdown files within a directory
func GetMarkdownFilePaths(dir string) ([]string, error) {
	var markdownFiles []string

	walkErr := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".md" {
			markdownFiles = append(markdownFiles, path)
		}

		return nil
	})
	if walkErr != nil {
		return nil, fmt.Errorf("Error walking directory %s: %v", dir, walkErr)
	}

	return markdownFiles, nil
}

func ParseMarkdownFile(filePath string) (ParsedMarkdownFile, error) {
	result := ParsedMarkdownFile{}

	// Read the markdown file
	content, readErr := os.ReadFile(filePath)
	if readErr != nil {
		return result, fmt.Errorf("Error reading file %s: %v", filePath, readErr)
	}

	// Create a new Goldmark parser
	md := goldmark.New(
		goldmark.WithExtensions(
			&frontmatter.Extender{},
		),
	)

	// Parse markdown file
	var buf bytes.Buffer
	ctx := parser.NewContext()
	if parseErr := md.Convert(content, &buf, parser.WithContext(ctx)); parseErr != nil {
		return result, fmt.Errorf("Error parsing markdown from %s: %v", filePath, parseErr)
	}

	// Get parsed frontmatter
	fm := frontmatter.Get(ctx)

	// Unmarshal the frontmatter into a struct
	var metadata Frontmatter
	if decodeErr := fm.Decode(&metadata); decodeErr != nil {
		return result, fmt.Errorf("Error decoding markdown from %s: %v", filePath, decodeErr)
	}

	// Get parsed Markdown content
	bufString := buf.String()
	parsedMarkdown := Unsafe(&bufString)

	result = ParsedMarkdownFile{
		Frontmatter: metadata,
		Html:        parsedMarkdown,
		Content:     &bufString,
	}

	return result, nil
}

func Unsafe(html *string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, *html)
		return
	})
}
