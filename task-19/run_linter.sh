#!/bin/sh
root_dir=$(pwd)
find . -name '*.go' -not -name '*_test.go' | while read file; do
	dir=$(dirname -- "$file")
	cd "$dir"
	golangci-lint run .
	cd "$root_dir"
done