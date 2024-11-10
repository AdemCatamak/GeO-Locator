package customErrors

type BasicCustomError struct {
	Message string
}

func NewBasicCustomError(message string) *BasicCustomError {
	return &BasicCustomError{
		Message: message,
	}
}

func (e BasicCustomError) Error() string {
	return e.Message
}
