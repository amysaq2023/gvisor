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

package stack

import (
	"sync"

	"gvisor.dev/gvisor/pkg/sentry/socket/plugin"
)

// Notifier holds all the state necessary to issue notifications when
// IO events occur on the observed FDs in plugin stack.
type Notifier struct {
	// the epoll FD used to register for io notifications.
	epFD int32

	// mu protects eventMap.
	mu sync.Mutex

	// eventMap maps file descriptors to their notification queues
	// and waiting status.
	eventMap map[uint32]*plugin.EventInfo
}

// AddFD implements plugin.PluginNotifier.AddFD.
func (n *Notifier) AddFD(fd uint32, eventInfo *plugin.EventInfo) error {
	//TODO: implement glue layer
	return nil
}

// RemoveFD implements plugin.PluginNotifier.RemoveFD.
func (n *Notifier) RemoveFD(fd uint32) {
	//TODO: implement glue layer
}

// UpdateFD implements plugin.PluginNotifier.UpdateFD.
func (n *Notifier) UpdateFD(fd uint32) error {
	//TODO: implement glue layer
	return nil
}

// notifier is a global variable that will be used by all sockets
// created in plugin network stack to report events.
var notifier = Notifier{}
