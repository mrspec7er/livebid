package trade

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"github.com/mrspec7er/livebid/server/internal/handler"
)

type Controller struct {
	Service  Service
	Response handler.ResponseJSON
	Upgrader websocket.Upgrader
}

func (c *Controller) Trade(w http.ResponseWriter, r *http.Request) {
	itemId := chi.URLParam(r, "itemId")

	wsConn := c.Service.WebsocketConnection()
	newConn, err := wsConn.Upgrade(w, r, nil)
	if err != nil {
		c.Response.GeneralErrorHandler(w, 500, err)
		return
	}

	c.Service.ProcessMessage(newConn, itemId)
}
