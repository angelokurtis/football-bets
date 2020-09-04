package http

const (
	UnknownResponse = iota
	InformationResponse
	SuccessfulResponse
	RedirectionResponse
	ClientErrorResponse
	ServerErrorResponse
)

type status struct {
	code int
}

func (s *status) Group() int {
	if s.code < 200 {
		return InformationResponse
	} else if s.code < 300 {
		return SuccessfulResponse
	} else if s.code < 400 {
		return RedirectionResponse
	} else if s.code < 500 {
		return ClientErrorResponse
	} else if s.code < 600 {
		return ServerErrorResponse
	} else {
		return UnknownResponse
	}
}
