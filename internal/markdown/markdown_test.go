package markdown

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestGetMarkdownFilePaths(t *testing.T) {
	// Create a temporary directory with sample markdown files
	tempDir := createTempDir(t)
	defer os.RemoveAll(tempDir)

	// Create sample markdown files in the temporary directory
	createTempFile(t, tempDir, "file1.md")
	createTempFile(t, tempDir, "file2.md")
	createTempFile(t, tempDir, "file3.txt")

	// Call the GetMarkdownFilePaths function
	filePaths, err := GetMarkdownFilePaths(tempDir)
	// Assert the results
	if err != nil {
		t.Fatalf("Error getting markdown files: %v", err)
	}

	expectedPaths := []string{
		filepath.Join(tempDir, "file1.md"),
		filepath.Join(tempDir, "file2.md"),
	}
	if !reflect.DeepEqual(filePaths, expectedPaths) {
		t.Errorf("Expected file paths: %v, got: %v", expectedPaths, filePaths)
	}
}

func TestParseMarkdownFile(t *testing.T) {
	tempFile := createTestFileForParsing(t)

	// Call the ParseMarkdownFile function
	result, err := ParseMarkdownFile(tempFile)
	// Assert the results
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedBody := "<h1>Heading 1</h1>\n<p>Sample content.</p>\n"
	if result.BodyAsString() != expectedBody {
		t.Errorf("Expected body: %q, got: %q", expectedBody, result.BodyAsString())
	}

	expectedFrontmatter := Frontmatter{
		Title:       "Sample Title",
		Description: "Sample Description",
		Date:        "2024-12-30",
	}
	if !reflect.DeepEqual(*result.Frontmatter(), expectedFrontmatter) {
		t.Errorf("Expected frontmatter: %+v, got: %+v", expectedFrontmatter, result.Frontmatter())
	}
}

// Helper functions for creating temporary files and directories
func createTempDir(t *testing.T) string {
	tempDir, err := os.MkdirTemp("", "markdown_test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	return tempDir
}

func createTempFile(t *testing.T, dir, name string) string {
	tempFile := filepath.Join(dir, name)
	_, err := os.Create(tempFile)
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	return tempFile
}

func createTestFileForParsing(t *testing.T) string {
	// Create a temporary markdown file with sample content
	tempFile := createTempFile(t, "", "sample.md")
	defer os.Remove(tempFile)

	// Write sample markdown content to the temporary file
	sampleContent := `---
title: Sample Title
description: Sample Description
date: 2024-12-30
---
# Heading 1
Sample content.
`
	err := os.WriteFile(tempFile, []byte(sampleContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write sample markdown file: %v", err)
	}
	return tempFile
}
