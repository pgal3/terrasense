package errors

type InvalidInputError struct {
	Message string
	Details map[string]any
}

func (e *InvalidInputError) Error() string {
	return e.Message
}
