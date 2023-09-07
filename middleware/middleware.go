package middleware

import (
	"net/http"
)

type MiddlewareFunc func(next HandlerFunc) HandlerFunc

type HandlerFunc func(c Context) error

type Context interface {
	// Request returns `*http.Request`.
	Request() *http.Request

	// String sends a string response with status code.
	String(code int, s string) error

	// JSON sends a JSON response with status code.
	JSON(code int, data any) error
}