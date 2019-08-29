#!/usr/bin/env bash
set -ex

abigen() {
  pushd abigen
  docker build -t abigen . && docker run -v /tmp/abigen-root:/root/out abigen "$@"
  popd
  cp /tmp/abigen-root/token.go .
  ls -la token.go
}

run() {
  docker-compose build
  docker-compose up
}

"$@"
