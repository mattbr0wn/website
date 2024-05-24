package markdown

import (
	"bytes"
	"context"
	"io"
	"os"

	"github.com/a-h/templ"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/frontmatter"
)

type Frontmatter struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Date        string `yaml:"date"`
	HeroImg     string `yaml:"hero"`
	Draft       bool   `yaml:"draft"`
}

type ArticleData struct {
	Metadata Frontmatter
	Body     *string
	Path     string
}

func ParseMarkdownFile(filePath string) (Frontmatter, templ.Component, *string) {

	// Read the markdown file
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
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
	if err := md.Convert(content, &buf, parser.WithContext(ctx)); err != nil {
		panic(err)
	}

	// Get parsed frontmatter
	fm := frontmatter.Get(ctx)

	// Unmarshal the frontmatter into a struct
	var metadata Frontmatter
	if err := fm.Decode(&metadata); err != nil {
		panic(err)
	}

	// Get parsed Markdown content
	bufString := buf.String()
	parsedMarkdown := Unsafe(&bufString)

	return metadata, parsedMarkdown, &bufString
}

func Unsafe(html *string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, *html)
		return
	})
}
