#!/bin/bash

set -e -o pipefail

version=${1-"latest"}

brwhte="$(tput setaf 15)"
bryllw="$(tput setaf 11)"
creset="$(tput sgr0)"
yarrow="${bryllw}=>${creset}"

build_binary() {
  set -u
  echo -e "${yarrow} building binary ${brwhte}testapp${creset}"
  docker pull casualjim/builder
  docker run --rm -it -v `pwd`:/go/src/github.com/casualjim/testapp -w /go/src/github.com/casualjim/testapp  casualjim/builder "$@"
}

build_container() {
  set -u
  echo -e "${yarrow} publishing ${brwhte}testapp${creset}"
  docker build --pull --no-cache -q -t casualjim/testapp:$version .
}

build_binary go get ./... && go build -o testapp -a -ldflags '-w -linkmode external -extldflags "-static"' .
build_container
