package ssg

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func deleteDirectory(dir_path string) error {
	if _, existErr := os.Stat(dir_path); os.IsNotExist(existErr) {
		return fmt.Errorf("Error removing directory %v: does not exist", dir_path)
	}
	err := os.RemoveAll(dir_path)
	if err != nil {
		return fmt.Errorf("Error removing directory: %v", err)
	}
	return nil
}

func createDirectory(dir_path string) error {
	err := os.MkdirAll(dir_path, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Error creating directory: %v", err)
	}
	return nil
}

func copyDirectoryContents(targetDirPath, destDirPath string) error {
	cmd := exec.Command("cp", "-r", targetDirPath, destDirPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Error copying directory from %s to %s: %v", targetDirPath, destDirPath, err)
	}
	return nil
}

func createFile(file_path string) (*os.File, error) {
	// Check if the file path is empty
	if file_path == "" {
		return nil, fmt.Errorf("Empty file path")
	}

	// Check if the directory of the file exists
	dir := filepath.Dir(file_path)
	if _, dirExistsErr := os.Stat(dir); os.IsNotExist(dirExistsErr) {
		return nil, fmt.Errorf("Directory %s does not exist", dir)
	}

	// Check if the file already exists
	if _, fileExistsErr := os.Stat(file_path); fileExistsErr == nil {
		return nil, fmt.Errorf("File %s already exists", file_path)
	}

	// Create the file
	f, createErr := os.Create(file_path)
	if createErr != nil {
		return nil, fmt.Errorf("Error creating file %s: %v", file_path, createErr)
	}
	return f, nil
}
