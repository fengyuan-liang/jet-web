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
	return api.Success(xlog.GenReqId(), r.CmdArgs), nil
}
```

