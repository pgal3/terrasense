package errors

type NotFoundError struct {
	Message string
	Details map[string]any
}

func (e *NotFoundError) Error() string {
	return e.Message
}
