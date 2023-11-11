// Copyright The Jet authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jet

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fengyuan-liang/jet-web/core/api"
	"github.com/fengyuan-liang/jet-web/core/router"
	"github.com/fengyuan-liang/jet-web/core/rpc"
	"github.com/fengyuan-liang/jet-web/core/sync"
	"github.com/fengyuan-liang/jet-web/pkg/xlog"
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
		xl := xlog.NewWith("jet-web")
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
	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds. quit := make(chan os.Signal, 1)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
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
