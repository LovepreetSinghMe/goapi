package handlers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/LovepreetSinghMe/goapi/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		jsonByte, _ := json.Marshal(map[string]string{"msg": "Not Found"})

		_ , err := w.Write([]byte(jsonByte))
		if err != nil {
			log.Error(err)
			return
		}
	})

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}
