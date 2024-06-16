package markdown

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/frontmatter"
)

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

func ParseMarkdownFile(filePath string) (*ParsedMarkdownFile, error) {
	result := &ParsedMarkdownFile{}

	fileContent, readErr := readMarkdownFile(filePath)
	if readErr != nil {
		return result, readErr
	}

	buf, ctx, parseErr := parseMarkdownBody(fileContent)
	if parseErr != nil {
		return result, parseErr
	}
	bufString := buf.String()
	result.SetBodyAsString(bufString)

	fm, fmErr := parseMarkdownFrontmatter(ctx)
	if fmErr != nil {
		return result, fmErr
	}
	result.SetFrontmatter(fm)

	return result, nil
}

func readMarkdownFile(filePath string) ([]byte, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error reading markdown file %s: %v", filePath, err)
	}
	return fileContent, nil
}

func parseMarkdownBody(fileContent []byte) (*bytes.Buffer, parser.Context, error) {
	var buf bytes.Buffer
	ctx := parser.NewContext()

	md := goldmark.New(goldmark.WithExtensions(&frontmatter.Extender{}))

	parseErr := md.Convert(fileContent, &buf, parser.WithContext(ctx))
	if parseErr != nil {
		return nil, nil, fmt.Errorf("Error parsing markdown body: %v", parseErr)
	}

	return &buf, ctx, nil
}

func parseMarkdownFrontmatter(ctx parser.Context) (Frontmatter, error) {
	fm := frontmatter.Get(ctx)

	var fileMetadata Frontmatter
	err := fm.Decode(&fileMetadata)
	if err != nil {
		return fileMetadata, fmt.Errorf("Error parsing markdown frontmatter: %v", err)
	}

	return fileMetadata, nil
}
