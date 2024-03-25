package errors

type InternalServerError struct {
	Message string
	Details map[string]any
}

func (e *InternalServerError) Error() string {
	return e.Message
}
