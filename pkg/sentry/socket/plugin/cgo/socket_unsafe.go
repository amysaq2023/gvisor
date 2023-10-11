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

package cgo

/*
#include <stdint.h>
#include <sys/epoll.h>
#include <sys/socket.h>

// epoll
int external_epoll_create(int size);
int external_epoll_ctl(int epfd, int op, int fd, struct epoll_event *event);
int external_epoll_wait(int epfd, struct epoll_event *events, int maxevents, int timeout);

// socket control operations
int external_socket(int domain, int type, int protocol);
int external_listen(int sockfd, int backlog);
int external_bind(int sockfd, const struct sockaddr *addr, socklen_t addrlen);
int external_accept(int sockfd, struct sockaddr *addr, socklen_t *addrlen, int flags);
int external_connect(int sockfd, const struct sockaddr *addr, socklen_t addrlen);
int external_getsockopt(int sockfd, int level, int optname,
			void *optval, socklen_t *optlen);
int external_setsockopt(int sockfd, int level, int optname,
			const void *optval, socklen_t optlen);
int external_getsockname(int sockfd, struct sockaddr *addr, socklen_t *addrlen);
int external_getpeername(int sockfd, struct sockaddr *addr, socklen_t *addrlen);
int external_ioctl(int d,  unsigned long int request, ...);
int external_shutdown(int sockfd, int how);
int external_close(int fd);

// receiving data operations
ssize_t external_recv(int sockfd, void *buf, size_t len, int flags);
ssize_t external_recvfrom(int sockfd, void *buf, size_t len, int flags,
			struct sockaddr *src_addr, socklen_t *addrlen);
ssize_t external_recvmsg(int sockfd, struct msghdr *msg, int flags);
ssize_t external_read(int fd, void *buf, size_t count);
ssize_t external_readv(int fd, const struct iovec *iov, int iovcnt);

// sending data operations
ssize_t external_send(int sockfd, const void *buf, size_t len, int flags);
ssize_t external_sendto(int sockfd, const void *buf, size_t len, int flags,
		const struct sockaddr *dest_addr, socklen_t addrlen);
ssize_t external_sendmsg(int sockfd, const struct msghdr *msg, int flags);
ssize_t external_write(int fd, const void *buf, size_t count);
ssize_t external_writev(int fd, const struct iovec *iov, int iovcnt);

*/
import "C"
import (
	"syscall"
)

// EpollCreate works as a CGO wrapper for external_epoll_create.
func EpollCreate() int {
	//TODO: implement cgo wrapper
	return 0
}

// EpollCTL works as a CGO wrapper for external_epoll_ctl.
func EpollCTL(epfd int32, op int, handle uint32, w uint64) {
	//TODO: implement cgo wrapper
}

// EpollWait works as a CGO wrapper for external_epoll_wait.
func EpollWait(epfd int32, events []syscall.EpollEvent, n int, us int) int {
	//TODO: implement cgo wrapper
	return 0
}

// Socket works as a CGO wrapper for external_socket.
func Socket(domain, skType, protocol int) int {
	//TODO: implement cgo wrapper
	return 0
}

// Bind works as a CGO wrapper for external_bind.
func Bind(handle uint32, sa []byte) int {
	//TODO: implement cgo wrapper
	return 0
}

// Listen works as a CGO wrapper for external_listen.
func Listen(handle uint32, backlog int) int {
	//TODO: implement cgo wrapper
	return 0
}

// Accept works as a CGO wrapper for external_accept.
func Accept(handle uint32, addrPtr *byte, lenPtr *uint32) int {
	//TODO: implement cgo wrapper
	return 0
}

// Ioctl works as a CGO wrapper for external_ioctl.
func Ioctl(handle uint32, cmd uint32, buf []byte) int {
	//TODO: implement cgo wrapper
	return 0
}

// Connect works as a CGO wrapper for external_connect.
func Connect(handle uint32, addr []byte) int {
	//TODO: implement cgo wrapper
	return 0
}

// Getsockopt works as a CGO wrapper for external_getsockopt.
func Getsockopt(handle uint32, l int, n int, val []byte, s int) (int, int) {
	//TODO: implement cgo wrapper
	return 0, 0
}

// Setsockopt works as a CGO wrapper for external_setsockopt.
func Setsockopt(handle uint32, l int, n int, val []byte) int {
	//TODO: implement cgo wrapper
	return 0
}

// Shutdown works as a CGO wrapper for external_shutdown.
func Shutdown(handle uint32, how int) int {
	//TODO: implement cgo wrapper
	return 0
}

// Close works as a CGO wrapper for external_close.
func Close(handle uint32) {
	//TODO: implement cgo wrapper
}

// Read works as a CGO wrapper for external_read.
func Read(handle uint32, buf uintptr, count int) int64 {
	//TODO: implement cgo wrapper
	return 0
}

// Readv works as a CGO wrapper for external_readv.
func Readv(handle uint32, iovs []syscall.Iovec) int64 {
	//TODO: implement cgo wrapper
	return 0
}

// Recvfrom works as a CGO wrapper for external_recvfrom.
func Recvfrom(handle uint32, buf, addr []byte, flags int) (int64, int) {
	//TODO: implement cgo wrapper
	return 0, 0
}

// Recvmsg works as a CGO wrapper for external_recvmsg.
func Recvmsg(handle uint32, iovs []syscall.Iovec, addr, control []byte, flags int) (int64, int, int, int) {
	//TODO: implement cgo wrapper
	return 0, 0, 0, 0
}

// Write works as a CGO wrapper for external_write.
func Write(handle uint32, buf uintptr, count int) int64 {
	//TODO: implement cgo wrapper
	return 0
}

// Writev works as a CGO wrapper for external_writev.
func Writev(handle uint32, iovs []syscall.Iovec) int64 {
	//TODO: implement cgo wrapper
	return 0
}

// Sendto works as a CGO wrapper for external_sendto.
func Sendto(handle uint32, buf uintptr, count int, flags int, addr []byte) int64 {
	//TODO: implement cgo wrapper
	return 0
}

// Sendmsg works as a CGO wrapper for external_sendmsg.
func Sendmsg(handle uint32, iovs []syscall.Iovec, addr []byte, flags int) int64 {
	//TODO: implement cgo wrapper
	return 0
}
