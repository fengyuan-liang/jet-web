// Copyright The Jet authors. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jet

// IJetController The IJetController interface is used to tag the routes you want to inject into Jet.
type IJetController interface {
	// JetControllerName
	// You can tag the name of the current router and easily retrieve the current router group during program execution.
	// You can also use Jet-defined aspects to handler common logic for the current router group.
	JetControllerName() string
}

type Controller struct{}

func (j *Controller) JetControllerName() string {
	return "jet_controller_test"
}
