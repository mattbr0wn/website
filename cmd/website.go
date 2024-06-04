package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/mattbr0wn/website/config"
	"github.com/mattbr0wn/website/internal/rss"
	"github.com/mattbr0wn/website/internal/ssg"
)

func main() {
	const port = "1616"
	const rootPath = "web/static"
	const buildCmd = "build"

	headData, err := config.HeadConfig()
	if err != nil {
		panic(err)
	}

	markdownFiles, err := ssg.GetMarkdownFiles("web/content")
	if err != nil {
		panic(err)
	}

	buildErr := buildStaticWebsite(rootPath, headData, &markdownFiles)
	if buildErr != nil {
		fmt.Printf("Error: Can't build website: %v", err)
	}
}

func buildStaticWebsite(rootPath string, headData config.HeadData, markdownFiles *[]string) error {
	fmt.Println("Setting up build...")
	// Remove the existing "static" directory
	if err := os.RemoveAll(rootPath); err != nil {
		fmt.Println("Error removing static directory:", err)
		return err
	}

	// Create the "static" and "static/img" directories
	if err := os.MkdirAll(filepath.Join(rootPath, "web/img"), os.ModePerm); err != nil {
		fmt.Println("Error creating img directory:", err)
		return err
	}

	// Copy image files into static
	cmd := exec.Command("cp", "-r", "web/img", rootPath)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error copying images to static:", err)
		return err
	}

	// Detect the current operating system and architecture
	goos := runtime.GOOS
	goarch := runtime.GOARCH

	templBinary := fmt.Sprintf("./bin/templ-%s-%s", goos, goarch)
	fmt.Println(templBinary)

	fmt.Println("Building static pages...")
	ssg.BuildStaticPages(rootPath, headData)
	fmt.Println("Building RSS feed...")
	rss.WriteRssFeed(rootPath, &headData, markdownFiles)
	fmt.Println("Build complete.")

	return nil
}
