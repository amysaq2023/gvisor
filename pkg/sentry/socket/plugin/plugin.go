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

// Package plugin provides a set of interfaces to interact with
// third-party netstack. It will be used during sandbox network setup when
// NetworkType is set as NetworkPlugin.
package plugin

import (
	"gvisor.dev/gvisor/pkg/sentry/inet"
	"gvisor.dev/gvisor/pkg/waiter"
)

// PluginStack defines a set of stack operations to work with a third-party
// plugin stack.
type PluginStack interface {
	inet.Stack

	// Init initializes plugin stack.
	Init(args *InitStackArgs) error

	// PreInit handles prepare steps before initializing plugin stack.
	// It may include joining namespace, mounting NIC, etc.
	PreInit(args *PreInitStackArgs) error

	// PostInit handles post steps after plugin stack initialized.
	PostInit(args *PostInitStackArgs) error

	// CleanUp removes plugin stack from sentry.
	// It will be used during destroying sandbox.
	CleanUp(args *CleanUpStackArgs) error
}

// InitStackArgs is a struct that holds arguments needed by PluginStack.Init.
type InitStackArgs struct {
	// TODO: implement arguments for Init
}

// PreInitStackArgs is a struct that holds arguments needed by
// PluginStack.PreInit.
type PreInitStackArgs struct {
	// TODO: implement arguments for PreInit
}

// PostInitStackArgs is a struct that holds arguments needed by
// PluginStack.PostInit.
type PostInitStackArgs struct {
	// TODO: implement arguments for PostInit
}

// CleanUpStackArgs is a struct that holds arguments needed by
// PluginStack.CleanUp.
type CleanUpStackArgs struct {
	// TODO: implement arguments for CleanUp
}

var pluginStack PluginStack

// RegisterPluginStack registers given stack as plugin stack.
func RegisterPluginStack(stack PluginStack) {
	if pluginStack != nil {
		panic("called RegisterPluginStack more than once")
	}
	pluginStack = stack
}

// GetPluginStack fetches the current registered plugin stack.
func GetPluginStack() PluginStack {
	return pluginStack
}

// EventInfo is a struct that holds information necessary to a socket
// notification mechanisms.
type EventInfo struct {
	// Queue is the socket corresponding event queue.
	wq *waiter.Queue

	// Mask represents events this socket registered.
	mask waiter.EventMask

	// Ready represents events has been currently reported.
	ready waiter.EventMask
}

// PluginNotifier represents a set of operations to handle
// plugin network stack's event notification mechanisms.
type PluginNotifier interface {
	// AddFD registers a new socket fd and its corresponding
	// event notification info into the global fdMap.
	AddFD(fd uint32, eventinfo *EventInfo) error

	// RemoveFD unregisters a socket fd and its corresponding
	// event notification info from the global fdMap.
	RemoveFD(fd uint32)

	// UpdateFD updates the set of events the socket fd needs
	// to be notified on.
	UpdateFD(fd uint32) error
}
