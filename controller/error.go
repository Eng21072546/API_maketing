package controller

func errorsToStrings(errs []error) []string {
	errorMessages := make([]string, len(errs))
	for i, err := range errs {
		errorMessages[i] = err.Error()
	}
	return errorMessages
}
