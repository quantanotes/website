package globals

import "net/http"

type Error struct {
	Status  int
	Message string
}

func NewError(status int, msg string) Error {
	return Error{status, msg}
}

func (e Error) Error() string {
	return e.Message
}

func ErrBadRequest(msg string) Error {
	if msg == "" {
		msg = "We couldn't process your request. Please check your input and try again."
	}
	return NewError(http.StatusBadRequest, msg)
}

func ErrUnauthorised(msg string) Error {
	if msg == "" {
		msg = "You do not have permission to perform this action. Please try again with the appropriate credentials."
	}
	return NewError(http.StatusUnauthorized, msg)
}
