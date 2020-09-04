package http

import "fmt"

type genericError struct {
	Method     string
	URL        string
	StatusCode int
}

type ClientError struct{ genericError }

func NewClientError(method string, URL string, statusCode int) *ClientError {
	return &ClientError{genericError{Method: method, URL: URL, StatusCode: statusCode}}
}

func (c *ClientError) Error() string {
	return fmt.Sprintf("the %s could not understand the %s request and answered %d", c.URL, c.Method, c.StatusCode)
}

type ServerError struct{ genericError }

func NewServerError(method string, URL string, statusCode int) *ServerError {
	return &ServerError{genericError{Method: method, URL: URL, StatusCode: statusCode}}
}

func (s *ServerError) Error() string {
	return fmt.Sprintf("the %s doesn't know how to handle the %s request and answered %d", s.URL, s.Method, s.StatusCode)
}
