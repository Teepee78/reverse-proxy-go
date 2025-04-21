package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"teepee78/reverse-proxy-go/config"
)

func ServeStatic(w http.ResponseWriter, r *http.Request) {
	if config.Vars.StaticDir == "" {
		return
	}

	path := r.URL.Path
	filePath := getFilePath(path)

	file, fileErr := os.Open(filePath)
	if fileErr != nil {
		fmt.Printf("Error opening file: %s", filePath)
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file", err)
		}
	}(file)

	setMimeType(w, filePath)

	_, copyErr := io.Copy(w, file)
	if copyErr != nil {
		fmt.Println("Error copying file", copyErr)
		return
	}
}
