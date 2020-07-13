# Getting Started

## Prerequisites

This list may not be complete.

- G++ v?
- JDK8
- Docker
- Docker Compose
- curl
- tar

## Setup

The directory of the `grchive-v2` document will be referenced to as `$GRCHIVE` in this document.
This document will walk you through setting up build environment and the necessary infrastructure to run the GRCHive app on your local machine.

1. Load the developer environment variables into your shell.

    ```
    source $GRCHIVE/config/dev_env
    ```
    
    Ensure your prompt has the `[GRC]` prefix. This script must be sourced before running any command.
1. Download the binaries needed for development.

    ```
    cd $GRCHIVE/deps/binaries
    ./download_all.sh
    ```
1. Add the following paths to your `$PATH`:

    ```
    $GRCHIVE/deps/binaries/bazel/output
    $GRCHIVE/deps/binaries/flyway
    ```
1. Create the Docker containers for the backend services:

    ```
    bazel run //devops/database:postgres
    ```
1. Run the Docker containers using Docker compose:

    ```
    cd $GRCHIVE/devops/docker
    docker-compose up
    ```
1. Apply the database migrations:

    ```
    cd $GRCHIVE/devops/database/vault
    ./migrate.sh
    ```
