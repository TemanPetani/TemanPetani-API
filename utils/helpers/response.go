package helpers

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func FailedResponse(message string) Response {
	return Response{
		Status:  "failed",
		Message: message,
	}
}

func SuccessResponse(message string) Response {
	return Response{
		Status:  "success",
		Message: message,
	}
}

func SuccessWithDataResponse(message string, data any) Response {
	return Response{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}
