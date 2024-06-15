package server

import (
	"net/http"
	"os"
	"time"

	"github.com/mrspec7er/livebid/server/internal/database"
)

type Config struct {
	DB *database.DBConn
}

func NewInstance(s *Config) *http.Server {

	server := &http.Server{
		Addr:         os.Getenv("PORT"),
		Handler:      s.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
