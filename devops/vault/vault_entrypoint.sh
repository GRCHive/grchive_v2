#!/usr/bin/dumb-init /bin/sh
function ctrl_c() {
    exit 0
}
trap ctrl_c SIGTERM SIGINT

# Ideally this apk add wouldn't be here but I don't think it's possible
# to do this within the Bazel file without some custom code so this is def easier.
apk add gettext
envsubst < /grchive/config.hcl > /vault/config/config.hcl

exec /usr/local/bin/docker-entrypoint.sh server
