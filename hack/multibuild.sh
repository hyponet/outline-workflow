#!/bin/sh

set -x
set -e

PKG=$1
BINDIR=$2

build() {
  echo "build OS=$1 Arch=$2"
  mkdir -p $BINDIR
  CGO_ENABLED=0 GOOS=$1 GOARCH=$2 go build -ldflags="-s -w" -o $BINDIR/outline-workflow-$1-$2 $PKG
}

main(){
  build "darwin" "amd64"
  build "darwin" "arm64"
}

main