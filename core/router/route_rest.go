// Copyright The jet authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package router

import (
	"jet-web/core/handler"
	"jet-web/core/rpc"
	"jet-web/pkg/xlog"
	"net/http"
	"reflect"
	"strings"
)

// ---------------------------------------------------------------------------

type Mux interface {
	Handle(pattern string, handler http.Handler)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
	SetDefault(handler http.Handler)
}

type Router struct {
	Factory       handler.HandlerFactory
	PatternPrefix string
	Separator     string
	Mux           Mux
	Default       http.Handler
}

func (r *Router) ListenAndServe(addr string, rcvr interface{}) error {

	return http.ListenAndServe(addr, r.Register(rcvr))
}

func (r *Router) Register(rcvr ...interface{}) Mux {

	if r.Mux == nil {
		r.Mux = rpc.NewServeMux()
	}
	if r.Default != nil {
		r.Mux.SetDefault(r.Default)
	}

	mux := r.Mux
	factory := r.Factory
	sep := r.Separator

	patternPrefix := r.PatternPrefix
	if strings.HasPrefix(patternPrefix, "/") {
		patternPrefix = patternPrefix[1:]
	}

	if factory == nil {
		factory = rpc.Factory
	}
	if sep == "" {
		sep = "_"
	}

	for _, rc := range rcvr {
		typ := reflect.TypeOf(rc)

		rcvr1 := reflect.ValueOf(rc)

		// Install the methods
		for m := 0; m < typ.NumMethod(); m++ {
			method := typ.Method(m)
			prefix, h, err := factory.Create(rcvr1, method)
			if err != nil {
				continue
			}
			pattern := []string{prefix}
			if patternPrefix != "" {
				pattern = append(pattern, patternPrefix)
			}
			pattern = append(pattern, patternOf(method.Name[len(prefix):], sep)...)

			mux.Handle(strings.Join(pattern, "/"), h)
			xlog.Debug("Install", pattern, "=>", method.Name)
		}
	}

	return mux
}

// AppleBanana => ["Apple", "Banana"]
// Apple_Banana => ["Apple", "*", "Banana"]
// AppleBanana_ => ["Apple", "Banana", "*"]
// Apple_Banana_ => ["Apple", "*", "Banana", "*"]
// ...
func patternOf(method string, sep string) (pattern []string) {

	for method != "" {
		pos := strings.Index(method, sep)
		if pos == -1 {
			return appendPattern(pattern, method)
		}
		if pos > 0 {
			pattern = appendPattern(pattern, method[:pos])
		}
		pattern = append(pattern, "*")
		method = method[pos+len(sep):]
	}
	return
}

func appendPattern(pattern []string, method string) []string {

	var i, last int
	for i = 1; i < len(method); i++ {
		c := method[i]
		if c >= 'A' && c <= 'Z' {
			pattern = append(pattern, method[last:i])
			last = i
		}
	}
	return append(pattern, method[last:i])
}

// ---------------------------------------------------------------------------
