// Copyright The Jet authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"jet-web/core/api"
	"jet-web/core/rpc"
	"jet-web/jet"
	"jet-web/pkg/xlog"
	"testing"
	"time"
)

func TestBoot(t *testing.T) {
	j := jet.NewWith(&UserController{}, &GoodsController{})
	j.StartService(":80")
}

type UserController struct{}

func (u *UserController) ApplicationName() string {
	return "UserController"
}

type Args struct {
	CmdArgs    []string
	FormParam1 string `json:"form_param1"`
	FormParam2 string `json:"form_param2"`
}

func (u *UserController) GetV1Usage0Week(r *Args, env *rpc.Env) (*api.Response, error) {
	time.Sleep(time.Millisecond * 10)
	return api.Success(xlog.GenReqId(), r.CmdArgs), nil
}

type GoodsController struct{}

func (u *GoodsController) ApplicationName() string {
	return "UserController"
}

func (u *GoodsController) GetV1Usage0Day(r *Args, env *rpc.Env) (*api.Response, error) {
	return api.Success(xlog.GenReqId(), r.CmdArgs), nil
}
