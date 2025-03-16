package main

import (
	"backend/pkg/logger"
	"backend/pkg/middleware"
	"fmt"
	"net/http"
)

func App() http.Handler {

	// config := configs.LoadConfig()
	// db := db.NewDb(config)
	router := http.NewServeMux()

	// Перенаправление на index.hmtl
	router.HandleFunc("/", middleware.Rewrite)

	// Repositories

	// Services

	// Handlers

	// Middlewars
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)
	return stack(router)
}

func main() {

	app := App()

	server := &http.Server{
		Addr:    ":5000",
		Handler: app,
	}

	fmt.Print("\033[H\033[2J")
	logger.INFO("http://localhost:5000")
	err := server.ListenAndServe()
	if err != nil {
		logger.ERROR(err)
	}
}
