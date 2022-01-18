package response

const (
	HeaderRequired = "1000"
	InternalError  = "5000"
)

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func OK(i interface{}) *Response {
	return &Response{
		Code:    "0",
		Message: "success",
		Data:    i,
	}
}

func Err(code, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
	}
}
