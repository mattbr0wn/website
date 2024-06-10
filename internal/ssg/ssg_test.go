package ssg

import (
	"os"
	"testing"
)

func TestCreateStaticHtmlPage(t *testing.T) {
	// Define the test cases
	testCases := []struct {
		staticURL   string
		expectedErr bool
	}{
		{"test.html", false},
		{"", true},
		{"invalid/path/test.html", true},
	}

	for _, tc := range testCases {
		// Call the function
		file, err := createStaticHtmlPage(tc.staticURL)

		if tc.expectedErr {
			// If an error is expected, check if the file is nil
			if err == nil {
				t.Errorf("Expected an error, but got a file for staticURL: %s", tc.staticURL)
			}
		} else {
			// If no error is expected, check if the file is created
			if file == nil {
				t.Errorf("Expected a file, but got an error for staticURL: %s", tc.staticURL)
			} else {
				// Close the file after the test
				file.Close()

				// Check if the file exists
				if _, err := os.Stat(tc.staticURL); os.IsNotExist(err) {
					t.Errorf("Expected file to be created, but it doesn't exist for staticURL: %s", tc.staticURL)
				}

				// Remove the file after the test
				os.Remove(tc.staticURL)
			}
		}
	}
}
