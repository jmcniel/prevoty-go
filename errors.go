package prevoty

type BadInputParameter struct{}

func (e *BadInputParameter) Error() string {
	return "Bad input parameter"
}

type BadAPIKey struct{}

func (e *BadAPIKey) Error() string {
	return "Bad API key"
}

type RequestTooLarge struct{}

func (e *RequestTooLarge) Error() string {
	return "Request too large"
}

type InternalError struct{}

func (e *InternalError) Error() string {
	return "Internal error"
}

type AccountQuotaExceeded struct{}

func (e *AccountQuotaExceeded) Error() string {
	return "Account quota exceeded"
}
