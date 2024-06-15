package ssg

import (
	"path/filepath"
	"testing"

	"github.com/mattbr0wn/website/config"
)

func TestGenerateStaticPath(t *testing.T) {
	testCases := []struct {
		name        string
		filePath    string
		expected    string
		expectError bool
	}{
		{
			name:     "Regular file",
			filePath: filepath.Join(config.CONTENT_DIR, "test.md"),
			expected: filepath.Join(config.ROOT_DIR, "test.html"),
		},
		{
			name:     "Index file",
			filePath: filepath.Join(config.CONTENT_DIR, "_index.md"),
			expected: filepath.Join(config.ROOT_DIR, "index.html"),
		},
		{
			name:     "File in subdirectory",
			filePath: filepath.Join(config.CONTENT_DIR, "sub", "test.md"),
			expected: filepath.Join(config.ROOT_DIR, "sub", "test.html"),
		},
		{
			name:        "File with different extension",
			filePath:    filepath.Join(config.CONTENT_DIR, "test.txt"),
			expected:    "",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := GenerateStaticPath(tc.filePath)
			if tc.expectError {
				if err == nil {
					t.Errorf("Expected an error, but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result != tc.expected {
					t.Errorf("Expected: %s, Got: %s", tc.expected, result)
				}
			}
		})
	}
}
