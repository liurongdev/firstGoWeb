package model

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func (res *Response) OK() *Response {
	res.Code = 200
	return res
}

func (res *Response) ERROR(code int) *Response {
	res.Code = code
	return res
}
