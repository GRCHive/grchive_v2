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

    container_pull(
        name = "fusionauth",
        registry = "index.docker.io",
        repository = "fusionauth/fusionauth-app",
        tag = "1.17.5"
    )

    container_pull(
        name = "vault",
        registry = "index.docker.io",
        repository = "library/vault",
        tag = "1.4.3"
    )
