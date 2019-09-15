#!/usr/bin/env bash

set -o errexit
set -o nounset
set -e
set -x

# =========================
# Runs Ginkgo/Gomega tests
# =========================

cd /tmp

echo "Installing Ginkgo and Gomega..."
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega/...

cd $GOPATH/src/github.com/Jaskaranbir/dapp-fs

echo "Running go test"
ginkgo --v -r -race
      # TODO: Integrate with codecov
      #  --cover \
      #  -coverprofile=coverage.txt \
      #  -outputdir=$(pwd)
