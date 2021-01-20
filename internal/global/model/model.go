package model

type TransportError struct {
	Err  error
	Code int
}

func (e *TransportError) Error() string {
	return e.Err.Error()
}

type Query struct {
	Query string `json:"query"`
}
