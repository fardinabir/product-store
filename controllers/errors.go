package controllers

import (
	"net/http"
)

type ErrorObj struct {
	Code    string
	Status  int
	Message string
}

func (e *ErrorObj) Error() string {
	return e.Message
}

var (
	ErrUserAlreadyExists   = &ErrorObj{Code: "UA409", Status: http.StatusConflict, Message: "User already exists"}
	ErrURINotFound         = &ErrorObj{Code: "UA404", Status: http.StatusNotFound, Message: "URL not found"}
	ErrMethodNotAllowed    = &ErrorObj{Code: "UA405", Status: http.StatusMethodNotAllowed, Message: "Method not allowed"}
	ErrInvalidData         = &ErrorObj{Code: "UA422", Status: http.StatusUnprocessableEntity, Message: "Invalid data"}
	ErrInternalServerError = &ErrorObj{Code: "UA500", Status: http.StatusInternalServerError, Message: "Internal server error"}
	ErrUnauthorizedReq     = &ErrorObj{Code: "UA401", Status: http.StatusUnauthorized, Message: "Unauthorized Access"}

	ErrNotFound = &ErrorObj{Code: "UA404", Status: http.StatusNotFound, Message: "Not found"}
)

func (e *ErrorObj) ErrorResponse() *Response {
	return &Response{Status: e.Status, Body: e}
}
