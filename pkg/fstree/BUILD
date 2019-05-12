load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["fstree.go"],
    importpath = "github.com/smukherj1/fs-tree/pkg/fstree",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["fstree_test.go"],
    embed = [":go_default_library"],
)
