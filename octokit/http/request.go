package http

import (
	"github.com/lostisland/go-sawyer"
)

type Headers map[string]string

type Request struct {
	Headers   Headers
	sawyerReq *sawyer.Request
}

func (r *Request) Head(output interface{}) (resp *Response, err error) {
	sawyerResp := r.sawyerReq.Head(output)
	if sawyerResp.IsError() {
		err = sawyerResp.ResponseError
		return
	}

	respErr := NewResponseError(sawyerResp)
	resp = &Response{Response: sawyerResp.Response, Error: respErr}

	return
}
