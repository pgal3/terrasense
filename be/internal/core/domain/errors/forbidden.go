package errors

type ForbiddenError struct {
	Message string `json:"message,omitempty"`
}

func (e *ForbiddenError) Error() string {
	return e.Message
}
