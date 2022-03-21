package message

type HttpResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func SuccessResponse(resData interface{}, message string) HttpResponse {
	return HttpResponse{Code: 1, Message: message, Result: resData}
}

func FailureResponse(resData interface{}, message string, code int) HttpResponse {
	return HttpResponse{Code: code, Message: message, Result: resData}
}
