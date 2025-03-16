package middleware

import (
	"net/http"
	"os"
	"path/filepath"
)

func Rewrite(w http.ResponseWriter, r *http.Request) {

	fileServer := http.FileServer(http.Dir("./ui"))
	// Проверяем, существует ли запрошенный файл
	filePath := filepath.Join("./ui", r.URL.Path)
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		// Если файл не существует, перенаправляем на index.html
		http.ServeFile(w, r, "./ui/index.html")
		return
	}

	// Если файл существует, отдаем его
	fileServer.ServeHTTP(w, r)
}
