package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mrspec7er/livebid/server/internal/module/item"
	"github.com/mrspec7er/livebid/server/internal/module/trade"
	"github.com/mrspec7er/livebid/server/internal/module/user"
)

func (s Config) RegisterRoutes() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello There!"})
	})

	router.Route("/users", user.Router(*s.DB))
	router.Route("/items", item.Router(*s.DB))
	router.Route("/trades", trade.Router(*s.DB))

	return router
}
