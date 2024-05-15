package errors

type UnknownError struct {
	Message       string         `json:"message,omitempty"`
	OriginalError string         `json:"original_error,omitempty"`
	Details       map[string]any `json:"details,omitempty"`
}

func (e *UnknownError) Error() string {
	return e.Message
}
