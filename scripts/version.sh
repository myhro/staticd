#!/usr/bin/env bash

set -eu

VERSION=$(git describe --all --dirty)

case $VERSION in
  tags/*)
    VERSION=${VERSION/tags\//}
    ;;
  heads/*)
    VERSION=$(git rev-parse --short HEAD)
    ;;&
  *-dirty)
    VERSION+="-dirty"
    ;;
esac

echo $VERSION
