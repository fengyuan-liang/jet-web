// Copyright The Jet authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"fmt"
	"github.com/fengyuan-liang/jet-web/core/api"
	"github.com/fengyuan-liang/jet-web/core/rpc"
	"github.com/fengyuan-liang/jet-web/jet"
	"github.com/fengyuan-liang/jet-web/pkg/xlog"
	"strconv"
	"testing"
	"time"
)

func TestBoot(t *testing.T) {
	j := jet.NewWith(&UserController{}, &GoodsController{})
	j.StartService(":8081")
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
	//time.Sleep(time.Millisecond * 10)
	return api.Success(xlog.GenReqId(), r.CmdArgs), nil
}

func (u *UserController) PostV1Usage0Month(env *rpc.Env) (*api.Response, error) {
	time.Sleep(time.Millisecond * 10)
	return api.Success(xlog.GenReqId(), fmt.Sprintf("month:%v", 111)), nil
}

type req struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (u *UserController) PostV1UsageMonth(r *req, env *rpc.Env) (*api.Response, error) {
	time.Sleep(time.Millisecond * 10)
	return api.Success(xlog.GenReqId(), fmt.Sprintf("month:%v", r)), nil
}

type GoodsController struct{}

func (u *GoodsController) ApplicationName() string {
	return "UserController"
}

func (u *GoodsController) PostV1Usage0Day(r *Args, env *rpc.Env) (*api.Response, error) {
	return api.Success(xlog.GenReqId(), r.CmdArgs), nil
}

func (u *GoodsController) GetV1Fib0(r *Args, env *rpc.Env) (*api.Response, error) {
	return api.Success(xlog.GenReqId(), Fibonacci(parseInt64(r.CmdArgs[0]))), nil
}

func Fibonacci(n int64) int64 {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-2) + Fibonacci(n-1)
}

func parseInt64(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("转换失败：", err)
		return 0
	}
	return num
}
