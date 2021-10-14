package limit

import (
	"github.com/go-chi/chi/v5"
	"github.com/nurislam03/golang_redis/data/repos"
)

// API provides application resources and handlers.
type API struct {
	Limit *LimitResource
}

// NewAPI configures and returns application API.
func NewAPI() (*API, error) {
	limitStore := repos.NewLimitStore()
	limit := NewLimitResource(limitStore)

	api := &API{
		Limit: limit,
	}
	return api, nil
}

// Router provides application routes.
func (a *API) Router() *chi.Mux {
	r := chi.NewRouter()

	r.Mount("/", a.Limit.router())

	return r
}
