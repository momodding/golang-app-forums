package response

type WebResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error,omitempty"`
}

func NewSuccessResponse(data interface{}, message string) *WebResponse {
	return &WebResponse{
		Code:    200,
		Message: message,
		Data:    data,
	}
}

func NewErrorResponse(errorCode int, error interface{}, message string) *WebResponse {
	return &WebResponse{
		Code:    errorCode,
		Message: message,
		Error:   error,
	}
}
