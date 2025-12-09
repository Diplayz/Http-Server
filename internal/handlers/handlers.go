package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func Register(mux *http.ServeMux) {
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/upload", UploadHandler)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	data, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Error loading page", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(data)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		file, header, err = r.FormFile("myFile")
	}

	if err != nil {
		http.Error(w, "Could not retrieve the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	if len(fileData) == 0 {
		http.Error(w, "File is empty", http.StatusInternalServerError)
		return
	}

	result, err := service.Convert(string(fileData))
	if err != nil {
		http.Error(w, "Error converting file", http.StatusInternalServerError)
		return
	}

	timeStr := time.Now().UTC().String()
	timeStr = strings.ReplaceAll(timeStr, ":", "-")
	timeStr = strings.ReplaceAll(timeStr, " ", "_")

	ext := filepath.Ext(header.Filename)
	if ext == "" {
		ext = ".txt"
	}

	filename := fmt.Sprintf("%s%s", timeStr, ext)

	createFile, err := os.Create(filename)
	if err != nil {
		http.Error(w, "Could not create file", http.StatusInternalServerError)
		return
	}
	defer createFile.Close()

	_, err = createFile.WriteString(result)
	if err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(result))
}
