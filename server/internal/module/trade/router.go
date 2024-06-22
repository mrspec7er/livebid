package trade

import (
	"github.com/go-chi/chi/v5"
	"github.com/mrspec7er/livebid/server/internal/database"
)

func Router(DBConn database.DBConn) func(chi.Router) {
	c := &Controller{
		Service: Service{
			Store: DBConn,
		},
	}

	return func(r chi.Router) {
		r.Get("/{itemId}", c.Trade)
	}
}
