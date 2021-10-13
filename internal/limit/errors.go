package limit

import (
	"net/http"

	"github.com/nurislam03/template/pkg/response"
)

var (
	ErrLimitIdCannotBeBlank = &response.ErrorCode{Code: "WLICBB400", Status: http.StatusBadRequest, Message: "Limit Id cannot be blank"}
	ErrLimitNotFound        = &response.ErrorCode{Code: "WLNF404", Status: http.StatusNotFound, Message: "Limit not found"}
)
