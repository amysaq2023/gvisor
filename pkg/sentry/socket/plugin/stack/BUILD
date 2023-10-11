package(licenses = ["notice"])

load("//tools:defs.bzl", "go_library")

go_library(
    name = "stack",
    srcs = [
        "notifier.go",
        "provider.go",
        "socket.go",
        "stack.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/abi/linux",
        "//pkg/abi/linux/errno",
        "//pkg/context",
        "//pkg/hostarch",
        "//pkg/marshal",
        "//pkg/sentry/arch",
        "//pkg/sentry/fsimpl/sockfs",
        "//pkg/sentry/inet",
        "//pkg/sentry/kernel",
        "//pkg/sentry/kernel/auth",
        "//pkg/sentry/kernel/time",
        "//pkg/sentry/socket",
        "//pkg/sentry/socket/plugin",
        "//pkg/sentry/socket/plugin/cgo",
        "//pkg/sentry/vfs",
        "//pkg/syserr",
        "//pkg/tcpip",
        "//pkg/tcpip/network/ipv4",
        "//pkg/tcpip/network/ipv6",
        "//pkg/usermem",
        "//pkg/waiter",
    ],
)
