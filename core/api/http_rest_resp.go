// Copyright The Jet authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"encoding/json"
	"net/http"
)

const CodeSuccess = 200

type Response struct {
	RequestId string `json:"request_id,omitempty"` //请求ID
	Code      int    `json:"code"`                 //错误码，200 成功，其他失败
	Message   string `json:"message,omitempty"`    //错误信息
	Data      any    `json:"data,omitempty"`
}

func (r *Response) Success() bool {
	return r.Code == CodeSuccess
}

func (r *Response) Error() string {
	data, _ := json.Marshal(r)
	return string(data)
}

func Success(reqId string, data any) *Response {
	return &Response{
		RequestId: reqId,
		Code:      CodeSuccess,
		Message:   "success",
		Data:      data,
	}
}

func Err(reqId string, code int, message string) *Response {
	return &Response{
		RequestId: reqId,
		Code:      code,
		Message:   message,
	}
}

func FromError(reqId string, err error) *Response {
	if apiErr, ok := err.(*Response); ok {
		return &Response{
			RequestId: reqId,
			Code:      apiErr.Code,
			Message:   apiErr.Message,
		}
	} else {
		return &Response{
			RequestId: reqId,
			Code:      http.StatusInternalServerError,
			Message:   err.Error(),
		}
	}
}
