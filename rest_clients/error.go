package clients

import (
	"fmt"
)

const (
	// success code
	CodeSuccess int64 = 0
	CodeClientError = 50010
	// error code
	RequestParameterInvalid int64 = 40001
	RequestDataExists       int64 = 40002
	RequestDataNotExisted   int64 = 40003
	AuthFailed              int64 = 40004
	PermissionDeny          int64 = 40005
	InternalError           int64 = 50000
	DatabaseError           int64 = 50001
	WorkflowError           int64 = 50002
	InternalServiceError    int64 = 50003
	ExternalServiceError    int64 = 50004
)

var codeMessageMap = map[int64]string{
	RequestParameterInvalid: "request parameter is invalid",
	RequestDataExists:       "request data already exists",
	RequestDataNotExisted:   "request data does not exists",
	AuthFailed:              "authorization failed",
	InternalError:           "internal server error",
	PermissionDeny:          "no permission",
	DatabaseError:           "database operation error",
	WorkflowError:           "workflow error",
	InternalServiceError:    "internal service error",
	ExternalServiceError:    "external service error",
}

func GetErrorMessage(code int64) (string, bool) {
	msg, ok := codeMessageMap[code]
	return msg, ok
}

type Error struct {
	Code    int64
	Message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("API Error: Code=%d, Message=%s", e.Code, e.Message)
}
