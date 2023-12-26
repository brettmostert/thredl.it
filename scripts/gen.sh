#!/usr/bin/env bash

echo "Code generation started (templ)"

CMD="templ generate"

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that directory
cd "${DIR}"


eval $CMD
echo "Code generation completed (templ)"
