package ssg

import (
	"fmt"
	"os"
	"os/exec"
)

func deleteDirectory(dir_path string) error {
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
	f, err := os.Create(file_path)
	if err != nil {
		return nil, fmt.Errorf("Error creating file %s: %v", file_path, err)
	}
	return f, nil
}
