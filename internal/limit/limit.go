package limit

import (
	"net/http"
	"strings"

	"github.com/nurislam03/golang_redis/data/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nurislam03/golang_redis/pkg/helpmate"
	"github.com/nurislam03/golang_redis/pkg/response"
)

type createLimitBodyPld struct {
	ServiceName         string `json:"service_name"`
	DailyLimit          int64  `json:"daily_limit"`
	DailyCountLimit     int64  `json:"daily_count_limit"`
	MonthlyLimit        int64  `json:"monthly_limit"`
	MonthlyCountLimit   int64  `json:"monthly_count_limit"`
	PerTransactionLimit int64  `json:"per_transaction_limit"`
}

func (body *createLimitBodyPld) Bind(_ *http.Request) error {
	body.ServiceName = strings.TrimSpace(body.ServiceName)
	return validation.ValidateStruct(body,
		validation.Field(&body.ServiceName, validation.Required),
		validation.Field(&body.DailyLimit, validation.Required),
		validation.Field(&body.DailyCountLimit, validation.Required),
		validation.Field(&body.MonthlyLimit, validation.Required),
		validation.Field(&body.MonthlyCountLimit, validation.Required),
		validation.Field(&body.PerTransactionLimit, validation.Required),
	)
}

func (rs *LimitResource) CreateLimit(w http.ResponseWriter, r *http.Request) {
	userSlug := helpmate.GetUserSlug(r)
	if userSlug == "" {
		_ = render.Render(w, r, response.ErrorResponse(response.ErrUserIDNotFound, ""))
		return
	}

	body := &createLimitBodyPld{}
	if err := render.Bind(r, body); err != nil {
		render.Render(w, r, response.ErrorResponse(response.ErrInvalidData, err))
		return
	}

	lt := &models.Limit{
		UserSlug:            userSlug,
		ServiceName:         body.ServiceName,
		DailyLimit:          body.DailyLimit,
		DailyCountLimit:     body.DailyCountLimit,
		MonthlyLimit:        body.MonthlyLimit,
		MonthlyCountLimit:   body.MonthlyCountLimit,
		PerTransactionLimit: body.PerTransactionLimit,
	}

	if err := rs.Store.Create(lt); err != nil {
		_ = render.Render(w, r, response.ErrorResponse(response.ErrInternalServer, err))
		return
	}

	render.Respond(w, r, response.Response{
		HTTPStatusCode: http.StatusOK,
		Data: response.Object{
			"status": "success",
		},
	})
}

func (rs *LimitResource) GetLimit(w http.ResponseWriter, r *http.Request) {
	userSlug := helpmate.GetUserSlug(r)
	if userSlug == "" {
		_ = render.Render(w, r, response.ErrorResponse(response.ErrUserIDNotFound, ""))
		return
	}
	ltID := strings.TrimSpace(chi.URLParam(r, "id"))
	if ltID == "" {
		_ = render.Render(w, r, response.ErrorResponse(ErrLimitIdCannotBeBlank, nil))
		return
	}

	lt, err := rs.Store.GetLimitByID(ltID)
	if err != nil {
		_ = render.Render(w, r, response.ErrorResponse(ErrLimitNotFound, nil))
		return
	}

	render.Respond(w, r, response.Response{
		HTTPStatusCode: http.StatusOK,
		Data: response.Object{
			"limit": lt,
		},
	})
}

func (rs *LimitResource) ListLimits(w http.ResponseWriter, r *http.Request) {
	userSlug := helpmate.GetUserSlug(r)
	if userSlug == "" {
		_ = render.Render(w, r, response.ErrorResponse(response.ErrUserIDNotFound, ""))
		return
	}

	qry := map[string]interface{}{}
	// qry["userSlug"] = userSlug
	sN := strings.TrimSpace(r.URL.Query().Get("serviceName"))
	if sN != "" {
		qry["serviceName"] = sN
	}

	ltList, err := rs.Store.GetLimitList(qry)
	if err != nil {
		_ = render.Render(w, r, response.ErrorResponse(response.ErrInternalServer, err))
	}

	render.Respond(w, r, response.Response{
		HTTPStatusCode: http.StatusOK,
		Data: response.Object{
			"list_limits": ltList,
		},
	})
}

type updateLimitBodyPld struct {
	ServiceName         string `json:"service_name"`
	DailyLimit          int64  `json:"daily_limit"`
	DailyCountLimit     int64  `json:"daily_count_limit"`
	MonthlyLimit        int64  `json:"monthly_limit"`
	MonthlyCountLimit   int64  `json:"monthly_count_limit"`
	PerTransactionLimit int64  `json:"per_transaction_limit"`
}

func (body *updateLimitBodyPld) Bind(_ *http.Request) error {
	body.ServiceName = strings.TrimSpace(body.ServiceName)
	return validation.ValidateStruct(body,
		validation.Field(&body.ServiceName, validation.Required),
		validation.Field(&body.DailyLimit, validation.Required),
		validation.Field(&body.DailyCountLimit, validation.Required),
		validation.Field(&body.MonthlyLimit, validation.Required),
		validation.Field(&body.MonthlyCountLimit, validation.Required),
		validation.Field(&body.PerTransactionLimit, validation.Required),
	)
}

func (rs *LimitResource) UpdateLimit(w http.ResponseWriter, r *http.Request) {
	userSlug := helpmate.GetUserSlug(r)
	if userSlug == "" {
		_ = render.Render(w, r, response.ErrorResponse(response.ErrUserIDNotFound, ""))
		return
	}

	body := &updateLimitBodyPld{}
	if err := render.Bind(r, body); err != nil {
		render.Render(w, r, response.ErrorResponse(response.ErrInvalidData, err))
		return
	}

	//id for specific limit entity
	ltID := strings.TrimSpace(chi.URLParam(r, "id"))
	if ltID == "" {
		_ = render.Render(w, r, response.ErrorResponse(ErrLimitIdCannotBeBlank, nil))
		return
	}

	//get limit by id
	lt, err := rs.Store.GetLimitByID(ltID)
	if err != nil {
		_ = render.Render(w, r, response.ErrorResponse(ErrLimitNotFound, nil))
		return
	}

	lt.ServiceName = body.ServiceName
	lt.DailyLimit = body.DailyLimit
	lt.DailyCountLimit = body.DailyCountLimit
	lt.MonthlyLimit = body.MonthlyLimit
	lt.MonthlyCountLimit = body.MonthlyCountLimit
	lt.PerTransactionLimit = body.PerTransactionLimit

	//update limit
	if err := rs.Store.Update(lt); err != nil {
		_ = render.Render(w, r, response.ErrorResponse(response.ErrInternalServer, nil))
		return
	}

	render.Respond(w, r, response.Response{
		HTTPStatusCode: http.StatusOK,
		Data: response.Object{
			"status": "limit updated successfully",
		},
	})
}
