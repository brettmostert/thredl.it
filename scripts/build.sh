#!/usr/bin/env bash
# Get the parent directory of where this script is.

echo Starting build...

# set defaults for flags
project="thredl.it"

# while getopts u:p:dbs: flag
while getopts p: flag
do
    case "${flag}" in
        p) project=${OPTARG};;
    esac
done

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that directory
cd "$DIR"

if [ ! -d $DIR/bin ]; then
    mkdir ./bin
fi

BASE_PRODUCT_VERSION=0.1
PROJECT_NAME=${project:="thredl.it"}
echo $PROJECT_NAME
BINARY_NAME=${PROJECT_NAME}
BIN_PATH=${BIN_PATH:=bin/${BINARY_NAME}}

# Get the git commit
GIT_COMMIT="$(git rev-parse HEAD)"
GIT_DIRTY="$(test -n "`git status --porcelain`" && echo "+CHANGES" || true)"

# Get the build date from the latest commit
function build_date() {
  : "${DATE_FORMAT:="%Y-%m-%dT%H:%M:%SZ"}"
  git show --no-show-signature -s --format=%cd --date=format:"$DATE_FORMAT" HEAD
}

BUILD_DATE=$(build_date)

cp -r ./public ./bin/
go build \
    -trimpath \
    -ldflags "
      -X 'github.com/brettmostert/thredl.it/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY}'
      -X 'github.com/brettmostert/thredl.it/version.Version=$BASE_PRODUCT_VERSION'
      -X 'github.com/brettmostert/thredl.it/version.BuildDate=$BUILD_DATE'
      " \
    -o "$BIN_PATH" \
    ./cmd/${project}

echo "Build completed SUCCESSFULLY"
