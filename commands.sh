#!/usr/bin/env bash
set -ex

abigen(){
  pushd abigen
  docker build -t abigen . && docker run abigen "$@"
}

runc() {
  docker-compose build
  docker-compose up
}

"$@"
