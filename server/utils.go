package server

import (
	"fmt"
	"mime"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"teepee78/reverse-proxy-go/config"
)

// PathUsageCache Maps a path to the used index
var PathUsageCache = make(map[string]int)

func getFilePath(path string) string {
	if path == "/" {
		return config.Vars.StaticDir + "/" + "index.html"
	}

	return getUrlPath(path)
}

func getUrlPath(path string) string {
	cleanedPath := cleanPath(path)
	return config.Vars.StaticDir + "/" + cleanedPath
}

func cleanPath(path string) string {
	if path[0] == '/' {
		return path[1:]
	}

	return path
}

func setMimeType(w http.ResponseWriter, filePath string) {
	extension := filepath.Ext(filePath)
	mimeType := mime.TypeByExtension(extension)

	if mimeType != "" {
		w.Header().Set("Content-Type", mimeType)
	}
}

func getUrl(targets []string, path string, uri string) *url.URL {
	target := getTarget(targets, path)
	urlString := target + strings.ReplaceAll(uri, path, "")
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return nil
	}

	return parsedUrl
}

// Round Robin implementation
func getTarget(targets []string, path string) string {
	index := getPathIndex(path)
	nextIndex := getNextIndex(index, len(targets))
	PathUsageCache[path] = nextIndex

	return targets[nextIndex]
}

func getPathIndex(path string) (index int) {
	pathIndex, exists := PathUsageCache[path]
	if !exists {
		index = 0
	} else {
		index = pathIndex
	}

	return
}

func getNextIndex(index int, length int) (next int) {
	next = index + 1

	if next > length-1 {
		next = 0
	}

	return
}
