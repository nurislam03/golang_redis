package response

import (
	"encoding/json"
	"github.com/go-chi/render"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

// APIError struct for response
type APIError struct {
	ID   string `json:"id"`
	Code string `json:"code,omitempty"`
	//Detail      json.RawMessage   `json:"detail,omitempty"`
	Errors         json.RawMessage `json:"errors,omitempty"`
	HTTPStatusCode int             `json:"network_code,omitempty"`
	StatusText     string          `json:"status,omitempty"`
	Title          string          `json:"title"`
	Source         error           `json:"source,omitempty"`
	tags           map[string]string
}

type (
	// Object represents as object
	Object map[string]interface{}
	// Errors represents as object
	Errors interface{}

	// Response represents a response
	Response struct {
		HTTPStatusCode int         `json:"network_code,omitempty"`
		Data           interface{} `json:"data,omitempty"`
		Meta           *pager      `json:"meta,omitempty"`
		Errors         []APIError  `json:"errors,omitempty"`
	}
)

// Render sets the application-specific error code in AppCode.
func (e *APIError) Render(_ http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func isJSON(s string) bool {
	var js interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

// ErrorResponse Global Error API response
func ErrorResponse(errorObj *ErrorCode, err interface{}) render.Renderer {
	var errorJson json.RawMessage

	switch errs := err.(type) {
	case error:
		if isJSON(errs.Error()) {
			errorJson = json.RawMessage(errs.Error())
		} else {
			errorJson, _ = json.Marshal(errs)
		}
	default:
		errorJson, _ = json.Marshal(errs)
	}

	return &APIError{
		ID:             uuid.NewV4().String(),
		HTTPStatusCode: errorObj.Status,
		Code:           errorObj.Code,
		Title:          errorObj.Message,
		Errors:         errorJson,
	}
}
