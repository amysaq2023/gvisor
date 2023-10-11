package(licenses = ["notice"])

load("//tools:defs.bzl", "go_library")

go_library(
    name = "plugin",
    srcs = [
        "plugin.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/sentry/inet",
        "//pkg/waiter",
    ],
)
