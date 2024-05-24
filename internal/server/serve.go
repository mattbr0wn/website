package server

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func ServeWebsite(port string, rootPath string) {
	log.Println("Server is running on http://localhost:" + port)

	http.HandleFunc("/", createPageHandler(rootPath))

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func createPageHandler(rootPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urlPath := r.URL.Path
		filePath := getFilePath(rootPath, urlPath)

		if filePath == "" {
			http.NotFound(w, r)
			return
		}

		if isDirectory(filePath) {
			serveDirectoryIndex(w, r, filePath)
			return
		}

		http.ServeFile(w, r, filePath)
	}
}

func serveDirectoryIndex(w http.ResponseWriter, r *http.Request, dirPath string) {
	indexFilePath := filepath.Join(dirPath, "index.html")
	if fileExists(indexFilePath) {
		http.ServeFile(w, r, indexFilePath)
	} else {
		http.NotFound(w, r)
	}
}

func getFilePath(rootPath string, urlPath string) string {
	filePath := filepath.Join(rootPath, urlPath)

	if isDirectory(filePath) {
		indexFilePath := filepath.Join(filePath, "index.html")
		if fileExists(indexFilePath) {
			return indexFilePath
		}
	} else {
		if fileExists(filePath) {
			return filePath
		}

		htmlFilePath := filePath + ".html"
		if fileExists(htmlFilePath) {
			return htmlFilePath
		}
	}

	return ""
}

func isDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
