package markdown

import (
	"strings"

	"github.com/mattbr0wn/website/config"
)

type Frontmatter struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Date        string `yaml:"date"`
	HeroImg     string `yaml:"hero"`
	Draft       bool   `yaml:"draft"`
}

type ParsedMarkdownFile struct {
	bodyAsString  string
	frontmatter   Frontmatter
	staticFileUrl string
}

func (p *ParsedMarkdownFile) BodyAsString() string {
	return p.bodyAsString
}

func (p *ParsedMarkdownFile) SetBodyAsString(body string) {
	p.bodyAsString = body
}

func (p *ParsedMarkdownFile) Frontmatter() Frontmatter {
	return p.frontmatter
}

func (p *ParsedMarkdownFile) SetFrontmatter(metadata Frontmatter) {
	p.frontmatter = metadata
}

func (p *ParsedMarkdownFile) StaticFileUrl() string {
	return p.staticFileUrl
}

func (p *ParsedMarkdownFile) SetStaticFileUrl(staticPath string) {
	p.staticFileUrl = strings.TrimPrefix(staticPath, config.ROOT_DIR)
}
