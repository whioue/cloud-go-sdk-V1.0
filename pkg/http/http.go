package http

type Resp struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	ErrCode int         `json:"errcode"`
	Data    interface{} `json:"data"`
}
