package errors

import "fmt"

var ErrOK = Error{
	Code:    "STATUS_OK",
	Message: "OK",
}

var ErrUnmarshalFail = Error{
	Code:    "SERVER_ERROR",
	Message: "unmarshal request failed",
}

type Error struct {
	Code    string
	Message string
}

func (e Error) Error() string {
	return fmt.Sprintf("Code: %s, Message: %s", e.Code, e.Message)
}
