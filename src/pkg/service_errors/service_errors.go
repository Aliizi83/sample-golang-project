package service_errors

type ServiceError struct {
	EndUserMessage string `json:"endUserMessage"`
	Err            error
}

func (e *ServiceError) Error() string {
	return e.EndUserMessage
}
