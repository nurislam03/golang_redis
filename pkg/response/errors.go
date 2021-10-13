package response

import "net/http"

type ErrorCode struct {
	Code    string
	Status  int
	Message string
}

// Error code style
// First character of Error message
// ErrUserNotFound = WUNF404, W = Service Name, U = User, N = Not , F = Found , 404 = HTTP Status Code
// ErrDataNotFound = WDNF404, W = Service Name, W = D = Data, N = Not, F = Found, 404 = HTTP Status Code
// ErrUnAuthorized = WUA401,  W = Service Name, UA = Unauthorized , 401 = HTTP Status Code
// ErrTokenExpired = WTE410,  W = Service Name, T = Token, E = Expired, 410 = HTTP Status Code
// ErrOTPResendExpired = WOTPRE410, W = Service Name, O = OTP, R=Resend, E = Expired, 410 = HTTP Status Code
// ErrOTPResendExpired = WORE410, W = Service Name, O = OTP, R=Resend, E = Expired, 410 = HTTP Status Code

var (
	ErrUserNotFound        = &ErrorCode{Code: "WUNF404", Status: http.StatusNotFound, Message: "User not found"}
	ErrUserIDNotFound      = &ErrorCode{Code: "WUIDNF404", Status: http.StatusNotFound, Message: "User ID not found"}
	ErrInvalidToken        = &ErrorCode{Code: "WIT422", Status: http.StatusUnprocessableEntity, Message: "Invalid token"}
	ErrTokenExpired        = &ErrorCode{Code: "WTE410", Status: http.StatusGone, Message: "Token is expired"}
	ErrURINotFound         = &ErrorCode{Code: "WURINF404", Status: http.StatusNotFound, Message: "URL not found"}
	ErrMethodNotAllowed    = &ErrorCode{Code: "WMNA405", Status: http.StatusMethodNotAllowed, Message: "Method not allowed"}
	ErrInvalidData         = &ErrorCode{Code: "WID422", Status: http.StatusUnprocessableEntity, Message: "Invalid data"}
	ErrPayloadTooLarge     = &ErrorCode{Code: "WPTL413", Status: http.StatusRequestEntityTooLarge, Message: "Payload size too large"}
	ErrTooManyRequest      = &ErrorCode{Code: "WTMR429", Status: http.StatusTooManyRequests, Message: "Too many request"}
	ErrInternalServerError = &ErrorCode{Code: "WISE500", Status: http.StatusInternalServerError, Message: "Internal server error"}
	ErrFromOtherService    = &ErrorCode{Code: "EFOS400", Status: http.StatusBadRequest, Message: "Bad request"}
	ErrUnAuthorized        = &ErrorCode{Code: "EUA401", Status: http.StatusUnauthorized, Message: "Unauthorized error"}
	ErrInvoiceNotFound     = &ErrorCode{Code: "EINF404", Status: http.StatusNotFound, Message: "Invoice not found"}
	ErrInternalServer      = &ErrorCode{Code: "EIS500", Status: http.StatusInternalServerError, Message: "Internal server error"}
	ErrOtpAlreadySent      = &ErrorCode{Code: "EOAS429", Status: http.StatusTooManyRequests, Message: "Otp already sent for this invoice"}
)
