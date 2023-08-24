// Copyright The Jet authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package xlog

import (
	"context"
	"net/http"
)

const (
	defaultCallDepth = 2
	logKey           = "X-Log"
	tagKey           = "X-Tag"
	uidKey           = "X-Uid"
	reqidKey         = "X-Reqid"
	billKey          = "X-Bill"
)

var UseReqCtx bool

func NewWithEnv(w http.ResponseWriter, req *http.Request) *Logger2 {

	reqId := req.Header.Get(reqidKey)
	if reqId == "" {
		reqId = genReqId()
		req.Header.Set(reqidKey, reqId)
	}
	h := w.Header()
	h.Set(reqidKey, reqId)

	var ctx context.Context
	if UseReqCtx {
		ctx = req.Context()
	}
	return &Logger2{h, reqId, defaultCallDepth, ctx}
}

// ============================================================================
// type *Logger

type Logger2 struct {
	h         http.Header
	reqId     string
	calldepth int //将 calldepth 暴露出来，默认值 defaultCallDepth = 2
	// https://jira.qiniu.io/browse/KODO-2916，将请求的 ctx 传递到下游
	// 为了兼容旧的接口，所以把 ctx 放到 logger 里面。
	// 其他地方不推荐使用这种方式，建议使用 xlog_context，将 xlog 放在 ctx 里面，使用 context 作为参数。
	ctx context.Context
}
