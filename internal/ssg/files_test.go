package ssg

import (
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestDeleteDirectory(t *testing.T) {
	tempDir, err := makeTestDirectory()
	defer os.RemoveAll(tempDir)

	// Test case 1: Deleting an existing directory
	err = deleteDirectory(tempDir)
	if err != nil {
		t.Errorf("Failed to delete directory: %v", err)
	}

	// Check if the directory no longer exists
	if _, err := os.Stat(tempDir); !os.IsNotExist(err) {
		t.Errorf("Directory still exists after deletion")
	}

	// Test case 2: Deleting a non-existent directory
	nonExistentDir := tempDir + "/nonexistent"
	err = deleteDirectory(nonExistentDir)
	if err == nil {
		t.Errorf("Expected an error when deleting a non-existent directory")
	}
}

func TestCreateDirectory(t *testing.T) {
	tempDir, err := makeTestDirectory()
	defer os.RemoveAll(tempDir)

	// Test case 1: Creating a new directory
	newDir := tempDir + "/newdirectory"
	err = createDirectory(newDir)
	if err != nil {
		t.Errorf("Failed to create directory: %v", err)
	}

	// Check if the directory exists
	if _, err := os.Stat(newDir); os.IsNotExist(err) {
		t.Errorf("Directory not created")
	}

	// Test case 2: Creating a directory that already exists
	err = createDirectory(newDir)
	if err != nil {
		t.Errorf("Unexpected error when creating an existing directory: %v", err)
	}

	// Test case 3: Creating a directory with nested directories
	nestedDir := tempDir + "/nested/directory"
	err = createDirectory(nestedDir)
	if err != nil {
		t.Errorf("Failed to create nested directory: %v", err)
	}

	// Check if the nested directory exists
	if _, err := os.Stat(nestedDir); os.IsNotExist(err) {
		t.Errorf("Nested directory not created")
	}
}

func TestCopyDirectoryContents(t *testing.T) {
	tempDir, err := makeTestDirectory()
	defer os.RemoveAll(tempDir)

	targetDir := filepath.Join(tempDir, "target")
	err = os.Mkdir(targetDir, os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to create target directory: %v", err)
	}

	// Create a file in the target directory
	fileContent := []byte("Hello, World!")
	filePath := filepath.Join(targetDir, "file.txt")
	err = os.WriteFile(filePath, fileContent, os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to create file in target directory: %v", err)
	}

	destDir := filepath.Join(tempDir, "dest")

	// Test case 1: Copy directory contents successfully
	err = copyDirectoryContents(targetDir, destDir)
	if err != nil {
		t.Errorf("Failed to copy directory contents: %v", err)
	}

	// Check if the file exists in the destination directory
	destFilePath := filepath.Join(destDir, "file.txt")
	if _, err := os.Stat(destFilePath); os.IsNotExist(err) {
		t.Errorf("File not copied to destination directory")
	}

	// Test case 2: Copy directory contents to a non-existent destination directory
	nonExistentDestDir := filepath.Join(tempDir, "nonexistent")
	err = copyDirectoryContents(targetDir, nonExistentDestDir)
	if err != nil {
		t.Errorf("Unexpected error when copying to a non-existent destination directory: %v", err)
	}

	// Check if the file exists in the non-existent destination directory
	nonExistentDestFilePath := filepath.Join(nonExistentDestDir, "file.txt")
	if _, err := os.Stat(nonExistentDestFilePath); os.IsNotExist(err) {
		t.Errorf("File not copied to non-existent destination directory")
	}

	// Test case 3: Copy from a non-existent source directory
	nonExistentSourceDir := filepath.Join(tempDir, "nonexistent_source")
	err = copyDirectoryContents(nonExistentSourceDir, destDir)
	if err == nil {
		t.Errorf("Expected an error when copying from a non-existent source directory")
	}
}

func TestCreateFile(t *testing.T) {
	tempDir, err := makeTestDirectory()
	defer os.RemoveAll(tempDir)

	// Test case 1: Create a new file successfully
	filePath := filepath.Join(tempDir, "file.txt")
	file, err := createFile(filePath)
	if err != nil {
		t.Errorf("Failed to create file: %v", err)
	}
	defer file.Close()

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		t.Errorf("File not created")
	}

	// Test case 2: Create a file in a non-existent directory
	nonExistentFilePath := filepath.Join(tempDir, "nonexistent", "file.txt")
	file, err = createFile(nonExistentFilePath)
	if err == nil {
		t.Errorf("Expected an error when creating a file in a non-existent directory")
	}
	if file != nil {
		file.Close()
	}

	// Test case 3: Create a file with an invalid path
	invalidFilePath := string([]byte{0x7f})
	file, err = createFile(invalidFilePath)
	if err == nil {
		t.Errorf("Expected an error when creating a file with an invalid path")
	}
	if file != nil {
		file.Close()
	}

	// Test case 4: Create a file with an empty path
	file, err = createFile("")
	if err == nil {
		t.Errorf("Expected an error when creating a file with an empty path")
	}
	if file != nil {
		file.Close()
	}

	// Test case 5: Create a file in an existing directory
	existingDirPath := filepath.Join(tempDir, "existing")
	err = os.Mkdir(existingDirPath, os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to create existing directory: %v", err)
	}
	existingFilePath := filepath.Join(existingDirPath, "file.txt")
	file, err = createFile(existingFilePath)
	if err != nil {
		t.Errorf("Failed to create file in an existing directory: %v", err)
	}
	defer file.Close()

	// Test case 6: Create a file that already exists
	file, err = createFile(existingFilePath)
	if err == nil {
		t.Errorf("Expected an error when creating a file that already exists")
	}
	if file != nil {
		file.Close()
	}
}

func makeTestDirectory() (string, error) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "test")
	if err != nil {
		log.Fatalf("Failed to create temporary directory: %v", err)
	}
	return tempDir, err
}
