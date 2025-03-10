package model

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func (response *Response) OK() *Response {
	response.Code = 200
	return response
}

func (response *Response) ERROR(code int) *Response {
	response.Code = code
	return response
}
