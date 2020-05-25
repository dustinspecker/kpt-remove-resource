#!/bin/bash
set -ex

if [ ! -f ./bin/kpt ]; then
  mkdir -p ./bin
  curl https://storage.googleapis.com/kpt-dev/releases/v0.24.0/linux_amd64/kpt_linux_amd64-v0.24.0.tar.gz |
  tar \
    --director ./bin \
    --extract \
    --gzip --file -
fi

./bin/kpt version
