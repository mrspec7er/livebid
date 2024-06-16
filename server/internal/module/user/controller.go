package user

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
	user := &database.User{}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		c.Response.BadRequestHandler(w)
		return
	}

	status, err := c.Service.Create(user)
	if err != nil {
		c.Response.GeneralErrorHandler(w, status, err)
		return
	}

	c.Response.MutationSuccessResponse(w, "User created successfully")
}

func (c *Controller) GetOne(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user := &database.User{}

	status, err := c.Service.FindOne(user, id)
	if err != nil {
		c.Response.GeneralErrorHandler(w, status, err)
		return
	}

	c.Response.QuerySuccessResponse(w, nil, user, nil)
}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user := &database.User{}

	status, err := c.Service.Delete(user, id)
	if err != nil {
		c.Response.GeneralErrorHandler(w, status, err)
		return
	}

	c.Response.MutationSuccessResponse(w, "User deleted successfully")
}
