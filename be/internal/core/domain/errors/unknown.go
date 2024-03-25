package errors

type UnknownError struct {
	Message string
	Details map[string]any
}

func (e *UnknownError) Error() string {
	return e.Message
}
