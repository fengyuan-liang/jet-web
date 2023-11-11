# Jet 🛩

一款和gin不太一样的golang web服务器

## usage

```go
func TestBoot(t *testing.T) {
	j := jet.NewWith(&UserController{})
	j.StartService(":80")
}
// 在Jet中 路由是挂载在Controller上的，通过Controller进行路由分组
type UserController struct{}
// 我们会尽可能的找到您需要的参数并将参数注入到您的结构体中
type Args struct {
	CmdArgs    []string
	FormParam1 string `json:"form_param1"`
	FormParam2 string `json:"form_param2"`
}

func (u *UserController) GetV1UsageWeek(r *Args, env *rpc.Env) (*api.Response, error) {
	return api.Success(xlog.GenReqId(), r.FormParam1), nil
}
```

我们注意到，`UserController`的方法比较有意思，叫`GetV1UsageWeek`，其实这代表着我们有一个接口`v1/usage/week`已经写好了，请求方式为`Get`，我们请求的参数会自动注入到`r *Args`中

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

## 更新计划

### 1. Hook

#### 1.1 参数相关

- [ ] 支持通过挂载hook对参数进行预解析、自定义参数校验规则
- [ ] 添加hook注入自定义的`context`，便于进行鉴权以及链路追踪等操作

### 2. 🤡Aspect（切面）支持

#### 2.1 常规切面

- [ ] 前置、后置、异常、环绕、最终五种切面

### 3. 路由策略

- [ ] 通过controller自定义路由前缀

### 4. 依赖注入支持

## Benchmarks

```shell
$ ab -c 400 -n 20000 http://localhost:8081/v1/usage/1111/week
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 2000 requests
Completed 4000 requests
Completed 6000 requests
Completed 8000 requests
Completed 10000 requests
Completed 12000 requests
Completed 14000 requests
Completed 16000 requests
Completed 18000 requests
Completed 20000 requests
Finished 20000 requests


Server Software:        JetServer
Server Hostname:        localhost
Server Port:            8081

Document Path:          /v1/usage/1111/week
Document Length:        76 bytes

Concurrency Level:      400
Time taken for tests:   1.661 seconds
Complete requests:      20000
Failed requests:        0
Total transferred:      4060000 bytes
HTML transferred:       1520000 bytes
Requests per second:    12041.08 [#/sec] (mean)
Time per request:       33.220 [ms] (mean)
Time per request:       0.083 [ms] (mean, across all concurrent requests)
Transfer rate:          2387.05 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.2      0       1
Processing:     8   33   2.4     33      39
Waiting:        1   17   8.8     17      37
Total:          8   33   2.4     33      39

Percentage of the requests served within a certain time (ms)
  50%     33
  66%     33
  75%     34
  80%     34
  90%     35
  95%     36
  98%     37
  99%     38
 100%     39 (longest request)
```

