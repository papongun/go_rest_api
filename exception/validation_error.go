package exception

type ValidationError struct {
	Message string
}

func (validationError ValidationError) Error() string {
	return validationError.Message
}

func (err ValidationError) parseErrorMessage() string {
	// TODO: Implement your custom error message here
	return "Field validation error"
}
