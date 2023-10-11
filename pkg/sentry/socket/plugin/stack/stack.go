// Copyright 2023 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package stack provides an implementation of plugin.PluginStack
// interface and an implementation of socket.Socket interface.
//
// It glues sentry interfaces with plugin netstack interfaces defined in cgo.
package stack

import (
	"gvisor.dev/gvisor/pkg/sentry/inet"
	"gvisor.dev/gvisor/pkg/sentry/socket/plugin"
)

// Stack is a struct that interacts with third-party network stack.
// It implements inet.Stack and plugin.PluginStack.
type Stack struct {
	inet.Stack
}

// Init implements plugin.PluginStack.Init.
func (s *Stack) Init(args *plugin.InitStackArgs) error {
	//TODO: implement glue layer
	return nil
}

// PreInit implements plugin.PluginStack.PreInit.
func (s *Stack) PreInit(args *plugin.PreInitStackArgs) error {
	//TODO: implement glue layer
	return nil
}

// PostInit implements plugin.PluginStack.PostInit.
func (s *Stack) PostInit(args *plugin.PostInitStackArgs) error {
	//TODO: implement glue layer
	return nil
}

// CleanUp implements plugin.PluginStack.CleanUp.
func (s *Stack) CleanUp(args *plugin.CleanUpStackArgs) error {
	//TODO: implement glue layer
	return nil
}

func init() {
	plugin.RegisterPluginStack(&Stack{})
}
