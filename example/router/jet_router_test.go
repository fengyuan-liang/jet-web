// Copyright The Jet authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package router

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
)

func init() {
	fmt.Printf("init ...")
}

type JetController struct{}

func (j *JetController) JetControllerName() string {
	return "jet_controller_test"
}

func getCurrentFunction() (fileName string, funcName string, line int) {
	pc, fileName, line, ok := runtime.Caller(1)
	if !ok {
		return "", "", 0
	}
	funcName = runtime.FuncForPC(pc).Name()
	return filepath.Base(fileName), filepath.Base(funcName), line
}

func TestBoot(t *testing.T) {
	// Get the current file name, function name and line number
	fileName, funcName, line := getCurrentFunction()

	// Output the results
	fmt.Printf("File: %s\nFunction: %s\nLine: %d\n", fileName, funcName, line)
}
