package main

import (
	"fmt"
	"net/http"

	"github.com/ChristianMe96/simple-go-api/internal/handlers"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	router := chi.NewRouter()

	handlers.Handler(router)

	fmt.Println("Starting GO API service....")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Error(err)
	}
}
