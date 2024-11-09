package customErrors

type CodedCustomError struct {
	BasicCustomError
	Code int
}

func NewCodedCustomError(message string, code int) *CodedCustomError {
	return &CodedCustomError{
		BasicCustomError: BasicCustomError{
			Message: message,
		},
		Code: code,
	}
}
