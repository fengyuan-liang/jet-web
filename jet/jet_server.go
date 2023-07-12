// Copyright The Jet authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jet

import (
	"jet-web/core/api"
	"jet-web/core/router"
	"jet-web/core/rpc"
	"jet-web/core/sync"
	"jet-web/pkg/xlog"
	"net/http"
)

type Service struct {
	httpServer *http.Server
	err        error
	stopD      sync.DoneChan
	router     router.Router
}

func New() *Service {
	xlog.SetOutputLevel(xlog.Ldebug)
	return &Service{
		stopD: sync.NewDoneChan(),
		router: router.Router{
			Separator: "0",
		},
		httpServer: &http.Server{},
	}
}

func NewWith(rcvr ...interface{}) *Service {
	s := New()
	s.Register(rcvr...)
	return s
}

func (s *Service) Register(rcvr ...interface{}) {
	s.httpServer.Handler = s.router.Register(rcvr...)
}

func (s *Service) StartService(addr string) {
	s.httpServer.Addr = addr
	go func() {
		xl := xlog.NewWith("[jet server]")
		xl.Infof("control server start on %s", addr)
		s.err = s.httpServer.ListenAndServe()
		s.stopD.SetDone()
	}()
	errCh := make(chan error)
	go func() {
		select {
		case <-s.StopD():
			xlog.Errorf("control service stopped, %+v", s.Error())
			errCh <- s.Error()
		}
	}()
	xlog.Errorf("exit with error %v", <-errCh)
}

func (s *Service) StopD() sync.DoneChanR {
	return s.stopD.R()
}

func (s *Service) SetStopD(stopD sync.DoneChan) {
	s.stopD = stopD
}

func (s *Service) SetRouter(r router.Router) {
	s.router = r
}

func (s *Service) SetHttpServer(h *http.Server) {
	s.httpServer = h
}

func (s *Service) Error() error {
	return s.err
}

func (s *Service) Stop() error {
	return s.httpServer.Close()
}

func (s *Service) GetStatus(env *rpc.Env) (*api.Response, error) {
	xl := xlog.NewWith("service")
	return api.Success(xl.ReqId, nil), nil
}
