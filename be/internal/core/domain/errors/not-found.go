package errors

type NotFoundError struct {
	Message string         `json:"message,omitempty"`
	Details map[string]any `json:"details,omitempty"`
}

func (e *NotFoundError) Error() string {
	return e.Message
}
