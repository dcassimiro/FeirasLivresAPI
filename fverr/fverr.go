package fverr

import (
	"errors"
	"fmt"
	"net/http"
)

// Error tipo de erro customizado
type Error struct {
	HTTPCode int         `json:"-"`
	Message  string      `json:"message"`
	Detail   interface{} `json:"detail,omitempty" swaggerignore:"true"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %v - message: %v - detail: %v", e.HTTPCode, e.Message, e.Detail)
}

// New cria um novo erro
func New(httpCode int, message string, detail interface{}) error {
	return &Error{
		HTTPCode: httpCode,
		Message:  message,
		Detail:   detail,
	}
}

// GetHTTPCode retorna o código http do erro
func GetHTTPCode(err error) int {
	var e *Error
	if !errors.As(err, &e) {
		return http.StatusInternalServerError
	}
	return e.HTTPCode
}

// NewError cria um novo struct Erro
func NewError(httpCode int, message string, detail interface{}) *Error {
	return &Error{
		HTTPCode: httpCode,
		Message:  message,
		Detail:   detail,
	}
}
