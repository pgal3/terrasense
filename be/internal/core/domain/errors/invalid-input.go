package errors

type InvalidInputError struct {
	Message string         `json:"message,omitempty"`
	Details map[string]any `json:"details,omitempty"`
}

func (e *InvalidInputError) Error() string {
	return e.Message
}
