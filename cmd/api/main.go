package main

import (
	"fmt"
	"net/http"

	"github.com/LovepreetSinghMe/goapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	log.SetLevel(log.TraceLevel)
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	var r = chi.NewRouter()
	handlers.Handler(r)
	fmt.Println("Starting GO API service")

	err := http.ListenAndServe("localhost:3000", r)
	if err != nil {
		log.Error(err)
	}
}
