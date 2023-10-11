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
	"gvisor.dev/gvisor/pkg/abi/linux"
	"gvisor.dev/gvisor/pkg/context"
	"gvisor.dev/gvisor/pkg/hostarch"
	"gvisor.dev/gvisor/pkg/marshal"
	"gvisor.dev/gvisor/pkg/sentry/arch"
	"gvisor.dev/gvisor/pkg/sentry/kernel"
	ktime "gvisor.dev/gvisor/pkg/sentry/kernel/time"
	"gvisor.dev/gvisor/pkg/sentry/socket"
	"gvisor.dev/gvisor/pkg/sentry/socket/plugin"
	"gvisor.dev/gvisor/pkg/sentry/vfs"
	"gvisor.dev/gvisor/pkg/syserr"
	"gvisor.dev/gvisor/pkg/usermem"
	"gvisor.dev/gvisor/pkg/waiter"
)

type socketOperations struct {
	vfsfd vfs.FileDescription
	vfs.FileDescriptionDefaultImpl
	vfs.DentryMetadataFileDescriptionImpl
	vfs.LockFD
	socket.SendReceiveTimeout

	family   int
	skType   linux.SockType
	protocol int

	// fd holds the current socket fd created by plugin stack.
	fd uint32 `state:"nosave"`

	// eventInfo holds current socket fd's corresponding notification queue
	// and events status. It will be used to interact with plugin
	// network stack for event reporting.
	eventInfo plugin.EventInfo
}

var _ = socket.Socket(&socketOperations{})

// Bind implements socket.Socket.Bind.
func (s *socketOperations) Bind(t *kernel.Task, sockaddr []byte) *syserr.Error {
	//TODO: implement glue layer
	return nil
}

// Listen implements socket.Socket.Listen.
func (s *socketOperations) Listen(t *kernel.Task, backlog int) *syserr.Error {
	//TODO: implement glue layer
	return nil
}

// Accept implements socket.Socket.Accept.
func (s *socketOperations) Accept(t *kernel.Task, peerRequested bool, flags int, blocking bool) (int32, linux.SockAddr, uint32, *syserr.Error) {
	//TODO: implement glue layer
	return 0, nil, 0, nil
}

// Connect implements socket.Socket.Connect.
func (s *socketOperations) Connect(t *kernel.Task, sockaddr []byte, blocking bool) *syserr.Error {
	//TODO: implement glue layer
	return nil
}

// Shutdown implements socket.Socket.Shutdown.
func (s *socketOperations) Shutdown(t *kernel.Task, how int) *syserr.Error {
	//TODO: implement glue layer
	return nil
}

// GetSockOpt implements socket.Socket.GetSockOpt.
func (s *socketOperations) GetSockOpt(t *kernel.Task, level int, name int, outPtr hostarch.Addr, outLen int) (marshal.Marshallable, *syserr.Error) {
	//TODO: implement glue layer
	return nil, nil
}

// SetSockOpt implements socket.Socket.SetSockOpt.
func (s *socketOperations) SetSockOpt(t *kernel.Task, level int, name int, optVal []byte) *syserr.Error {
	//TODO: implement glue layer
	return nil
}

// RecvMsg implements socket.Socket.RecvMsg.
func (s *socketOperations) RecvMsg(t *kernel.Task, dst usermem.IOSequence, flags int, haveDeadline bool, deadline ktime.Time, senderRequested bool, controlDataLen uint64) (int, int, linux.SockAddr, uint32, socket.ControlMessages, *syserr.Error) {
	//TODO: implement glue layer
	return 0, 0, nil, 0, socket.ControlMessages{}, syserr.ErrInvalidArgument
}

// SendMsg implements socket.Socket.SendMsg.
func (s *socketOperations) SendMsg(t *kernel.Task, src usermem.IOSequence, to []byte, flags int, haveDeadline bool, deadline ktime.Time, controlMessages socket.ControlMessages) (int, *syserr.Error) {
	//TODO: implement glue layer
	return 0, nil
}

// State implements socket.Socket.State.
func (s *socketOperations) State() uint32 {
	//TODO: implement glue layer
	return 0
}

// Type implements socket.Socket.Type.
func (s *socketOperations) Type() (family int, skType linux.SockType, protocol int) {
	//TODO: implement glue layer
	return 0, 0, 0
}

// Close implements socket.Socket.Close.
func (s *socketOperations) OnClose(ctx context.Context) error {
	return nil
}

// EventRegister implements socket.Socket.EventRegister.
func (s *socketOperations) EventRegister(e *waiter.Entry) error {
	//TODO: implement glue layer
	return nil
}

// EventUnregister implements socket.Socket.EventUnregister.
func (s *socketOperations) EventUnregister(e *waiter.Entry) {
	//TODO: implement glue layer
}

// Readiness implements socket.Socket.Readiness.
func (s *socketOperations) Readiness(mask waiter.EventMask) waiter.EventMask {
	//TODO: implement glue layer
	return 0
}

// Read implements socket.Socket.Read.
func (s *socketOperations) Read(ctx context.Context, dst usermem.IOSequence, opts vfs.ReadOptions) (int64, error) {
	//TODO: implement glue layer
	return 0, nil
}

// Write implements socket.Socket.Write.
func (s *socketOperations) Write(ctx context.Context, src usermem.IOSequence, opts vfs.WriteOptions) (int64, error) {
	//TODO: implement glue layer
	return 0, nil
}

// Epollable implements socket.Socket.Epollable.
func (s *socketOperations) Epollable() bool {
	return true
}

// Ioctl implements socket.Socket.Ioctl.
func (s *socketOperations) Ioctl(ctx context.Context, io usermem.IO, sysno uintptr, args arch.SyscallArguments) (uintptr, error) {
	//TODO: implement glue layer
	return 0, nil
}

// Release implements socket.Socket.Release.
func (s *socketOperations) Release(ctx context.Context) {
	//TODO: implement glue layer
}

// GetSockName implements socket.Socket.GetSockName.
func (s *socketOperations) GetSockName(t *kernel.Task) (linux.SockAddr, uint32, *syserr.Error) {
	//TODO: implement glue layer
	return nil, 0, nil
}

// GetPeerName implements socket.Socket.GetPeerName.
func (s *socketOperations) GetPeerName(t *kernel.Task) (linux.SockAddr, uint32, *syserr.Error) {
	//TODO: implement glue layer
	return nil, 0, nil
}
