package errors

type InternalServerError struct {
	Message       string         `json:"message,omitempty"`
	OriginalError string         `json:"original_error,omitempty"`
	Details       map[string]any `json:"details,omitempty"`
}

func (e *InternalServerError) Error() string {
	return e.Message
}
