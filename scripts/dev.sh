#!/usr/bin/env bash

echo "Starting app"

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that directory
cd "${DIR}/bin"

BINARY_NAME="./thredl.it"
eval $BINARY_NAME
