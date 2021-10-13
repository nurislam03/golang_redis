package limit

import "github.com/go-chi/chi/v5"

// LimitResource ...
type LimitResource struct {
	Store LimitStore
}

// NewLimitResource ...
func NewLimitResource(store LimitStore) *LimitResource {
	return &LimitResource{
		Store: store,
	}
}

func (rs *LimitResource) router() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", rs.CreateLimit)
	r.Get("/{id}", rs.GetLimit)
	r.Get("/", rs.ListLimits)
	r.Put("/{id}", rs.UpdateLimit)
	return r
}
