package errors

import "net/http"

type appErrorCode int

const (
	StatusBadRequest appErrorCode = iota + 1
	StatusUnAuthorized
	StatusPaymentRequired
	StatusForbidden
	StatusNotFound
	StatusMethodNotAllowed
	StatusNotAcceptable
	StatusProxyAuthRequired
	StatusRequestTimeout
	StatusConflict
	StatusGone
	StatusLengthRequired
	StatusPreconditionFailed
	StatusRequestEntityTooLarge
	StatusRequestURITooLong
	StatusUnsupportedMediaType
	StatusRequestedRangeNotSatisfiable
	StatusExpectationFailed
	StatusTeapot
	StatusMisdirectedRequest
	StatusUnprocessableEntity
	StatusLocked
	StatusFailedDependency
	StatusTooEarly
	StatusUpgradeRequired
	StatusPreconditionRequired
	StatusTooManyRequests
	StatusRequestHeaderFieldsTooLarge
	StatusUnavailableForLegalReasons
)

func errorCode(code appErrorCode) string {
	var codex map[appErrorCode]string = map[appErrorCode]string{
		StatusBadRequest:                   "IE-01",
		StatusUnAuthorized:                 "IE-02",
		StatusPaymentRequired:              "IE-03",
		StatusForbidden:                    "IE-04",
		StatusNotFound:                     "IE-05",
		StatusMethodNotAllowed:             "IE-06",
		StatusNotAcceptable:                "IE-07",
		StatusProxyAuthRequired:            "IE-08",
		StatusRequestTimeout:               "IE-09",
		StatusConflict:                     "IE-10",
		StatusGone:                         "IE-11",
		StatusLengthRequired:               "IE-12",
		StatusPreconditionFailed:           "IE-13",
		StatusRequestEntityTooLarge:        "IE-14",
		StatusRequestURITooLong:            "IE-15",
		StatusUnsupportedMediaType:         "IE-16",
		StatusRequestedRangeNotSatisfiable: "IE-17",
		StatusExpectationFailed:            "IE-18",
		StatusTeapot:                       "IE-19",
		StatusMisdirectedRequest:           "IE-20",
		StatusUnprocessableEntity:          "IE-21",
		StatusLocked:                       "IE-22",
		StatusFailedDependency:             "IE-23",
		StatusTooEarly:                     "IE-24",
		StatusUpgradeRequired:              "IE-25",
		StatusPreconditionRequired:         "IE-26",
		StatusTooManyRequests:              "IE-27",
		StatusRequestHeaderFieldsTooLarge:  "IE-28",
		StatusUnavailableForLegalReasons:   "IE-29",
	}
	return codex[code]
}

var (
	BadRequest                   = &ApplicationError{Code: StatusBadRequest, HttpStatusCode: http.StatusBadRequest, ErrorCode: errorCode(StatusBadRequest), Message: message(StatusBadRequest)}
	UnAuthorized                 = &ApplicationError{Code: StatusUnAuthorized, HttpStatusCode: http.StatusUnauthorized, ErrorCode: errorCode(StatusUnAuthorized), Message: message(StatusUnAuthorized)}
	PaymentRequired              = &ApplicationError{Code: StatusPaymentRequired, HttpStatusCode: http.StatusPaymentRequired, ErrorCode: errorCode(StatusPaymentRequired), Message: message(StatusPaymentRequired)}
	Forbidden                    = &ApplicationError{Code: StatusForbidden, HttpStatusCode: http.StatusForbidden, ErrorCode: errorCode(StatusForbidden), Message: message(StatusForbidden)}
	NotFound                     = &ApplicationError{Code: StatusNotFound, HttpStatusCode: http.StatusNotFound, ErrorCode: errorCode(StatusNotFound), Message: message(StatusNotFound)}
	MethodNotAllowed             = &ApplicationError{Code: StatusMethodNotAllowed, HttpStatusCode: http.StatusMethodNotAllowed, ErrorCode: errorCode(StatusMethodNotAllowed), Message: message(StatusMethodNotAllowed)}
	NotAcceptable                = &ApplicationError{Code: StatusNotAcceptable, HttpStatusCode: http.StatusNotAcceptable, ErrorCode: errorCode(StatusNotAcceptable), Message: message(StatusNotAcceptable)}
	ProxyAuthRequired            = &ApplicationError{Code: StatusProxyAuthRequired, HttpStatusCode: http.StatusProxyAuthRequired, ErrorCode: errorCode(StatusProxyAuthRequired), Message: message(StatusProxyAuthRequired)}
	RequestTimeout               = &ApplicationError{Code: StatusRequestTimeout, HttpStatusCode: http.StatusRequestTimeout, ErrorCode: errorCode(StatusRequestTimeout), Message: message(StatusRequestTimeout)}
	Conflict                     = &ApplicationError{Code: StatusConflict, HttpStatusCode: http.StatusConflict, ErrorCode: errorCode(StatusConflict), Message: message(StatusConflict)}
	Gone                         = &ApplicationError{Code: StatusGone, HttpStatusCode: http.StatusGone, ErrorCode: errorCode(StatusGone), Message: message(StatusGone)}
	LengthRequired               = &ApplicationError{Code: StatusLengthRequired, HttpStatusCode: http.StatusLengthRequired, ErrorCode: errorCode(StatusLengthRequired), Message: message(StatusLengthRequired)}
	PreconditionFailed           = &ApplicationError{Code: StatusPreconditionFailed, HttpStatusCode: http.StatusPreconditionFailed, ErrorCode: errorCode(StatusPreconditionFailed), Message: message(StatusPreconditionFailed)}
	RequestEntityTooLarge        = &ApplicationError{Code: StatusRequestEntityTooLarge, HttpStatusCode: http.StatusRequestEntityTooLarge, ErrorCode: errorCode(StatusRequestEntityTooLarge), Message: message(StatusRequestEntityTooLarge)}
	RequestURITooLong            = &ApplicationError{Code: StatusRequestURITooLong, HttpStatusCode: http.StatusRequestURITooLong, ErrorCode: errorCode(StatusRequestURITooLong), Message: message(StatusRequestURITooLong)}
	UnsupportedMediaType         = &ApplicationError{Code: StatusUnsupportedMediaType, HttpStatusCode: http.StatusUnsupportedMediaType, ErrorCode: errorCode(StatusUnsupportedMediaType), Message: message(StatusUnsupportedMediaType)}
	RequestedRangeNotSatisfiable = &ApplicationError{Code: StatusRequestedRangeNotSatisfiable, HttpStatusCode: http.StatusRequestedRangeNotSatisfiable, ErrorCode: errorCode(StatusRequestedRangeNotSatisfiable), Message: message(StatusRequestedRangeNotSatisfiable)}
	ExpectationFailed            = &ApplicationError{Code: StatusExpectationFailed, HttpStatusCode: http.StatusExpectationFailed, ErrorCode: errorCode(StatusExpectationFailed), Message: message(StatusExpectationFailed)}
	Teapot                       = &ApplicationError{Code: StatusTeapot, HttpStatusCode: http.StatusTeapot, ErrorCode: errorCode(StatusTeapot), Message: message(StatusTeapot)}
	MisdirectedRequest           = &ApplicationError{Code: StatusMisdirectedRequest, HttpStatusCode: http.StatusMisdirectedRequest, ErrorCode: errorCode(StatusMisdirectedRequest), Message: message(StatusMisdirectedRequest)}
	UnprocessableEntity          = &ApplicationError{Code: StatusUnprocessableEntity, HttpStatusCode: http.StatusUnprocessableEntity, ErrorCode: errorCode(StatusUnprocessableEntity), Message: message(StatusUnprocessableEntity)}
	Locked                       = &ApplicationError{Code: StatusLocked, HttpStatusCode: http.StatusLocked, ErrorCode: errorCode(StatusLocked), Message: message(StatusLocked)}
	FailedDependency             = &ApplicationError{Code: StatusFailedDependency, HttpStatusCode: http.StatusFailedDependency, ErrorCode: errorCode(StatusFailedDependency), Message: message(StatusFailedDependency)}
	TooEarly                     = &ApplicationError{Code: StatusTooEarly, HttpStatusCode: http.StatusTooEarly, ErrorCode: errorCode(StatusTooEarly), Message: message(StatusTooEarly)}
	UpgradeRequired              = &ApplicationError{Code: StatusUpgradeRequired, HttpStatusCode: http.StatusUpgradeRequired, ErrorCode: errorCode(StatusUpgradeRequired), Message: message(StatusUpgradeRequired)}
	PreconditionRequired         = &ApplicationError{Code: StatusPreconditionRequired, HttpStatusCode: http.StatusPreconditionRequired, ErrorCode: errorCode(StatusPreconditionRequired), Message: message(StatusPreconditionRequired)}
	TooManyRequests              = &ApplicationError{Code: StatusTooManyRequests, HttpStatusCode: http.StatusTooManyRequests, ErrorCode: errorCode(StatusTooManyRequests), Message: message(StatusTooManyRequests)}
	RequestHeaderFieldsTooLarge  = &ApplicationError{Code: StatusRequestHeaderFieldsTooLarge, HttpStatusCode: http.StatusRequestHeaderFieldsTooLarge, ErrorCode: errorCode(StatusRequestHeaderFieldsTooLarge), Message: message(StatusRequestHeaderFieldsTooLarge)}
	UnavailableForLegalReasons   = &ApplicationError{Code: StatusUnavailableForLegalReasons, HttpStatusCode: http.StatusUnavailableForLegalReasons, ErrorCode: errorCode(StatusUnavailableForLegalReasons), Message: message(StatusUnavailableForLegalReasons)}
)

func getApplicationError(code appErrorCode) *ApplicationError {
	switch code {
	case StatusBadRequest:
		return BadRequest
	case StatusUnAuthorized:
		return UnAuthorized
	case StatusPaymentRequired:
		return PaymentRequired
	case StatusForbidden:
		return Forbidden
	case StatusNotFound:
		return NotFound
	case StatusMethodNotAllowed:
		return MethodNotAllowed
	case StatusNotAcceptable:
		return NotAcceptable
	case StatusProxyAuthRequired:
		return ProxyAuthRequired
	case StatusRequestTimeout:
		return RequestTimeout
	case StatusConflict:
		return Conflict
	case StatusGone:
		return Gone
	case StatusLengthRequired:
		return LengthRequired
	case StatusPreconditionFailed:
		return PreconditionFailed
	case StatusRequestEntityTooLarge:
		return RequestEntityTooLarge
	case StatusRequestURITooLong:
		return RequestURITooLong
	case StatusUnsupportedMediaType:
		return UnsupportedMediaType
	case StatusRequestedRangeNotSatisfiable:
		return RequestedRangeNotSatisfiable
	case StatusExpectationFailed:
		return ExpectationFailed
	case StatusTeapot:
		return Teapot
	case StatusMisdirectedRequest:
		return MisdirectedRequest
	case StatusUnprocessableEntity:
		return UnprocessableEntity
	case StatusLocked:
		return Locked
	case StatusFailedDependency:
		return FailedDependency
	case StatusTooEarly:
		return TooEarly
	case StatusUpgradeRequired:
		return UpgradeRequired
	case StatusPreconditionRequired:
		return PreconditionRequired
	case StatusTooManyRequests:
		return TooManyRequests
	case StatusRequestHeaderFieldsTooLarge:
		return RequestHeaderFieldsTooLarge
	case StatusUnavailableForLegalReasons:
		return UnavailableForLegalReasons
	}
	return nil
}

func message(code appErrorCode) string {
	switch code {
	case StatusBadRequest:
		return "Bad Request"
	case StatusUnAuthorized:
		return "Not Authorized"
	case StatusPaymentRequired:
		return "Payment Required"
	case StatusForbidden:
		return "Forbidden"
	case StatusNotFound:
		return "Not Found"
	case StatusMethodNotAllowed:
		return "Method Not Allowed"
	case StatusNotAcceptable:
		return "Not Acceptable"
	case StatusProxyAuthRequired:
		return "Proxy Auth Required"
	case StatusRequestTimeout:
		return "Request Timeout"
	case StatusConflict:
		return "Conflict"
	case StatusGone:
		return "Gone"
	case StatusLengthRequired:
		return "Length Required"
	case StatusPreconditionFailed:
		return "Precondition Failed"
	case StatusRequestEntityTooLarge:
		return "Request Entity Too Large"
	case StatusRequestURITooLong:
		return "Request URI Too Long"
	case StatusUnsupportedMediaType:
		return "Unsupported Media Type"
	case StatusRequestedRangeNotSatisfiable:
		return "Requested Range Not Satisfiable"
	case StatusExpectationFailed:
		return "Expectation Failed"
	case StatusTeapot:
		return "Teapot"
	case StatusMisdirectedRequest:
		return "Misdirected Request"
	case StatusUnprocessableEntity:
		return "Unprocessable Entity"
	case StatusLocked:
		return "Locked"
	case StatusFailedDependency:
		return "Failed Dependency"
	case StatusTooEarly:
		return "Too Early"
	case StatusUpgradeRequired:
		return "Upgrade Required"
	case StatusPreconditionRequired:
		return "Precondition Required"
	case StatusTooManyRequests:
		return "Too Many Requests"
	case StatusRequestHeaderFieldsTooLarge:
		return "Request Header Fields TooLarge"
	case StatusUnavailableForLegalReasons:
		return "Unavailable For Legal Reasons"
	}
	return ""
}
