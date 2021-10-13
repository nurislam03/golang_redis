package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"

	"github.com/nurislam03/template/internal/limit"
	"github.com/nurislam03/template/pkg/dbconn"
	"github.com/nurislam03/template/pkg/logging"
	"github.com/nurislam03/template/pkg/response"
)

// New configures application resources and routes.
func New() (*chi.Mux, error) {

	router, logger := InitAndBindRouter()

	err := dbconn.Connect()
	if err != nil {
		logger.WithField("module", "mongodb connect").Error(err)
		return nil, err
	}

	limitAPI, err := limit.NewAPI()
	if err != nil {
		logger.WithField("module", "limit").Error(err)
		return nil, err
	}

	router.Group(func(rt chi.Router) {
		rt.Mount("/api/v1/template/limits", limitAPI.Router())
	})

	return router, nil
}

func InitAndBindRouter() (*chi.Mux, *logrus.Logger) {
	logger := logging.NewLogger()

	router := chi.NewRouter()
	router.Use(Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	//router.Use(middleware.DefaultCompress)
	router.Use(middleware.Timeout(15 * time.Second))
	router.Use(middleware.Heartbeat("/ping"))

	router.Use(logging.NewStructuredLogger(logger))
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		_ = render.Render(w, r, response.ErrorResponse(response.ErrURINotFound, nil))
	})

	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		_ = render.Render(w, r, response.ErrorResponse(response.ErrMethodNotAllowed, nil))
	})

	return router, logger
}
