#!/usr/bin/env bash

HASH="$(git rev-parse --short HEAD)"
NAME="asciidoc-link-verifier"
VERSION="0.1"
DATE=$(date +%d-%m-%Y" "%H:%M:%S)
BUILD_DIR="build"
BUILD_TMP_DIR="$BUILD_DIR/tmp/$NAME-$VERSION"
BUILD_BIN_DIR="$BUILD_DIR/binaries"
BUILD_NIX_BIN_FILE="$BUILD_TMP_DIR/$NAME"
BUILD_WIN_BIN_FILE="$BUILD_TMP_DIR/$NAME.exe"

rm -rf ${BUILD_DIR}
mkdir -p ${BUILD_TMP_DIR}
mkdir -p ${BUILD_BIN_DIR}
cp LICENSE ${BUILD_TMP_DIR}
cp README.adoc ${BUILD_TMP_DIR}

echo "Merging third-party license files..."
sh third-party-licenses.sh
cp "$BUILD_DIR/tmp/THIRD-PARTY-LICENSES" ${BUILD_TMP_DIR}

echo "Generating binaries..."

nixBinary() {
    echo "➡ $1/$2"
    GOOS=$1 GOARCH=$2 go build -ldflags "-s -w -X main.Version=$VERSION -X main.CommitHash=$HASH -X 'main.CompileDate=$DATE'" -o ${BUILD_NIX_BIN_FILE}
    tar -czf "$BUILD_BIN_DIR/$NAME-$VERSION-$3.tar.gz" -C ${BUILD_TMP_DIR} .
    rm -f ${BUILD_NIX_BIN_FILE}
}

winBinary() {
    echo "➡ $1/$2"
    GOOS=$1 GOARCH=$2 go build -ldflags "-s -w -X main.Version=$VERSION -X main.CommitHash=$HASH -X 'main.CompileDate=$DATE'" -o ${BUILD_WIN_BIN_FILE}
    zip -r -q -T -j "$BUILD_BIN_DIR/$NAME-$VERSION-$3.zip" ${BUILD_TMP_DIR}
    rm -f ${BUILD_WIN_BIN_FILE}
}

# Mac
nixBinary darwin amd64 osx

# NetBSD
nixBinary netbsd amd64 netbsd64
nixBinary netbsd 386 netbsd32

# OpenBSD
nixBinary openbsd amd64 openbsd64
nixBinary openbsd 386 openbsd32

# FreeBSD
nixBinary freebsd amd64 freebsd64
nixBinary freebsd 386 freebsd32

# Linux
nixBinary linux amd64 linux64
nixBinary linux 386 linux32
nixBinary linux arm arm

# Windows
winBinary windows amd64 win64
winBinary windows 386 win32

rm -rf ${BUILD_TMP_DIR}