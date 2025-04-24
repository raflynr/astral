package helper

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Errors  any    `json:"errors,omitempty"`
}

func NewSuccess(detail any) *Response {
	return &Response{
		Message: "success",
		Data:    detail,
	}
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

func NewError(message string, detail any) error {
	return &ErrorResponse{
		Message: message,
		Errors:  detail,
	}
}
