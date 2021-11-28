package api

import (
	"github.com/nurislam03/golang_redis/config"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

// API ...
type API struct {
	router chi.Router
	cfg    *config.Config
}

// NewAPI ...
func NewAPI(cfg *config.Config) *API {
	api := &API{
		router: chi.NewRouter(),
		cfg:    cfg,
	}

	api.register()
	api.registerRouter()
	return api
}

var logger = logrus.New()

func init() {
	logger.SetLevel(logrus.DebugLevel)
}

// Handler ...
func (api *API) Handler() http.Handler {
	return api.router
}

func (api *API) register() {

	api.router.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: logger}))
	api.router.Use(recoverer)

	api.router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		err := newAPIError("Not Found", errURINotFound, nil)
		panic(err)
	})

	api.router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		err := newAPIError("Method Not Allowed", errInvalidMethod, nil)
		resp := response{
			code:   http.StatusMethodNotAllowed,
			Errors: []apiError{*err},
		}
		resp.serveJSON(w)
	})
}

func (api *API) registerRouter() {
	api.router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/system", api.systemHandlers())
		r.Mount("/home", api.homeHandlers())
	})
}

func (api *API) homeHandlers() http.Handler {
	h := chi.NewRouter()
	h.Group(func(r chi.Router) {
		r.Get("/", api.Home)
	})

	return h
}

func (api *API) systemHandlers() http.Handler {
	h := chi.NewRouter()
	h.Group(func(r chi.Router) {
		r.Get("/check", api.systemCheck)
		r.Get("/panic", api.systemPanic)
		r.Get("/err", api.systemErr)
		r.Get("/verr", api.systemValidationErr)
	})

	return h
}
