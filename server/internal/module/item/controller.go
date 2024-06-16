package item

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mrspec7er/livebid/server/internal/database"
	"github.com/mrspec7er/livebid/server/internal/handler"
)

type Controller struct {
	Service  Service
	Response handler.ResponseJSON
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	item := &database.Item{}

	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		c.Response.BadRequestHandler(w)
		return
	}

	status, err := c.Service.Create(item)
	if err != nil {
		c.Response.GeneralErrorHandler(w, status, err)
		return
	}

	c.Response.MutationSuccessResponse(w, "Item created successfully")
}

func (c *Controller) GetOne(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	item := &database.Item{}

	status, err := c.Service.FindOne(item, id)
	if err != nil {
		c.Response.GeneralErrorHandler(w, status, err)
		return
	}

	c.Response.QuerySuccessResponse(w, nil, item, nil)
}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	item := &database.Item{}

	status, err := c.Service.Delete(item, id)
	if err != nil {
		c.Response.GeneralErrorHandler(w, status, err)
		return
	}

	c.Response.MutationSuccessResponse(w, "Item deleted successfully")
}
