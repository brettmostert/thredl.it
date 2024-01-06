#!/usr/bin/env bash
# Get the parent directory of where this script is.

echo Linting...

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that directory
cd "$DIR"

golangci-lint run
nilaway ./...
