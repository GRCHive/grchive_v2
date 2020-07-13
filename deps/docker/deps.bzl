load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_pull",
)

def load_docker_deps():
    container_pull(
        name = "postgres",
        registry = "index.docker.io",
        repository = "library/postgres",
        tag = "12.3"
    )
