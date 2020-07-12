workspace(
    name = "grchive_v2",
    managed_directories = {"@npm": ["src/node_modules"]},
)

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# NODE
http_archive(
    name = "build_bazel_rules_nodejs",
    sha256 = "cb6d92c93a1769205d6573c21363bdbdcf5831af114a7fbc3f800b8598207dee",
    urls = ["https://github.com/bazelbuild/rules_nodejs/releases/download/2.0.0-rc.2/rules_nodejs-2.0.0-rc.2.tar.gz"],
)

load("@build_bazel_rules_nodejs//:index.bzl", "node_repositories", "npm_install")

node_repositories(
    package_json = ["//src:package.json"],
    node_version = "14.4.0",
    node_repositories = {
        "14.4.0-linux_amd64": ("node-v14.4.0-linux-x64.tar.gz", "node-v14.4.0-linux-x64", "8e219f15f496d975910c3964d7ccb7b88d4dc68992b52a18396e05280b1cd642"),
    },
    node_urls = ["https://nodejs.org/dist/v{version}/{filename}"],
)

npm_install(
    name = "npm",
    package_json = "//src:package.json",
    package_lock_json = "//src:package-lock.json"
)

# GOLANG
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "a8d6b1b354d371a646d2f7927319974e0f9e52f73a2452d2b3877118169eb6bb",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.23.3/rules_go-v0.23.3.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.23.3/rules_go-v0.23.3.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

http_archive(
    name = "bazel_gazelle",
    sha256 = "cdb02a887a7187ea4d5a27452311a75ed8637379a1287d8eeb952138ea485f7d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.21.1/bazel-gazelle-v0.21.1.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.21.1/bazel-gazelle-v0.21.1.tar.gz",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

load("//deps/golang:deps.bzl", "load_golang_deps")

# gazelle:repository_macro deps/golang/deps.bzl%load_golang_deps
load_golang_deps()
