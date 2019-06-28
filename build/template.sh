#!/bin/bash
set -eu
BASEDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
RAWURL=https://raw.githubusercontent.com/mamemomonga/notebook-go/master/build

usage() {
	echo "USAGE: $0 app_path"
	exit 1
}

download() {
	local p=$1
	local d
	if [ -n "${2:-}" ]; then d=$2; else d=$1; fi
	echo " Download $p -> $d"
	curl -sL $RAWURL/$p > $d
}

if [ -z "${1:-}" ]; then usage; fi
APPPATH=$1
APPNAME=$(basename $APPPATH)

mkdir -v $APPNAME
cd $APPNAME

go mod init $APPPATH
download Dockerfile
download Makefile
perl -i -npE 's#sampleapp#'$APPNAME'#sg' Makefile

download build.mk
mkdir -p $APPNAME/buildinfo
download sampleapp/buildinfo/buildinfo.go $APPNAME/buildinfo/buildinfo.go

echo '0.0.0' > version
echo '0' > revision

mkdir -p $APPNAME
cat > $APPNAME/main.go << EOS
package main

import (
	"fmt"

	"$APPPATH/$APPNAME/buildinfo"
)

func main() {
	fmt.Printf("Version: %s Revision: %s\n",buildinfo.Version, buildinfo.Revision)
}
EOS
make
bin/$APPNAME
