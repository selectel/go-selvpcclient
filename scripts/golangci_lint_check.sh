#!/usr/bin/env bash

echo "==> Running golangci-lint..."
golangci-lint run \
    --enable golint \
    --enable interfacer \
    --enable unconvert \
    --enable goconst \
    --enable gocyclo \
    --enable gofmt \
    --enable goimports \
    --enable maligned \
    --enable depguard \
    --enable unparam \
    --enable nakedret \
    --enable prealloc \
    ./...
if [ $? -eq 1 ]; then
    echo ""
    echo "Golangci-lint found suspicious constructs. Please check the reported"; \
    echo "constructs and fix them if necessary before submitting the code for review."; \
    exit 1
fi

exit 0