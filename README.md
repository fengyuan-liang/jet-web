# jet

一款和gin不太一样的golang web服务器

## usage

```go
func TestBoot(t *testing.T) {
	j := jet.NewWith(&UserController{})
	j.StartService(":80")
}

type UserController struct{}

type Args struct {
	CmdArgs    []string
	FormParam1 string `json:"form_param1"`
	FormParam2 string `json:"form_param2"`
}

func (u *UserController) GetV1UsageWeek(r *Args, env *rpc.Env) (*api.Response, error) {
	return api.Success(xlog.GenReqId(), r.FormParam1), nil
}
```

我们注意到，`UserController`的方法比较有意思，叫`GetV1UsageWeek`，其实这代表着我们有一个接口`v1/usage/week`已经写好了，请问方式为`Get`，我们请求的参数会自动注入到`r *Args`中

```shell
$ curl http://localhost/v1/usage/week?form_param1=1
{"request_id":"ZRgQg3Osptrx","code":200,"message":"success","data":"1"}
```

如果想要定义`v1/usage/week/1`的形式，或者`v1/usage/1/week`，我们可以使用`0`或其他符号填充名字

```go
GetV1UsageWeek0 -> v1/usage/week/1 // 0的位置表示要接受一个可变的参数
GetV1Usage0Week -> v1/usage/1/week
```

参数会默认注入到`CmdArgs`中

```go
func (u *UserController) GetV1Usage0Week(r *Args, env *rpc.Env) (*api.Response, error) {
	return api.Success(xlog.GenReqId(), r.CmdArgs), nil
}
```

```shell
$ curl http://localhost/v1/usage/1/week
{"request_id":"H5OQ4Jg0yBtg","code":200,"message":"success","data":["1"]}
```

