package handlers

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/s444v/go-sixth-sprint/internal/service"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	content, err := os.ReadFile("../index.html")
	if err != nil {
		http.Error(w, "Cant read file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	MainHandler(w, r)
	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	tmpFile, err := os.OpenFile(time.Now().UTC().Format("2006-01-02_15-04-05")+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer tmpFile.Close()
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if _, err := tmpFile.WriteString(line); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "%s", service.Detector(line))
	}
}
