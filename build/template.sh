#!/bin/bash
# ひな形から作成

set -eu
BASEDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
RAWURL=https://raw.githubusercontent.com/mamemomonga/notebook-go/master/build

usage() {
	echo "USAGE: $0 import_path"
	exit 1
}

download() {
	local p=$1
	local d
	if [ -n "${2:-}" ]; then d=$2; else d=$1; fi
	echo " Download $p -> $d"
	curl -sL $RAWURL/$p > $d
}

# download() {
# 	local p=$1
# 	local d
# 	if [ -n "${2:-}" ]; then d=$2; else d=$1; fi
# 	echo " Copy $p -> $d"
# 	cp $BASEDIR/$p $d
# }

if [ -z "${1:-}" ]; then usage; fi
APPPATH=$1
APPNAME=$(basename $APPPATH)

mkdir -v $APPNAME
cd $APPNAME

go mod init $APPPATH
mkdir vendor

download Dockerfile
download Makefile
download build.mk
perl -i -npE 's#sampleapp#'$APPNAME'#sg' Makefile

download build.mk
mkdir -p $APPNAME/buildinfo
download sampleapp/buildinfo/buildinfo.go $APPNAME/buildinfo/buildinfo.go

echo '0.0.0' > version

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

cat > .gitignore << 'EOS'
/revision
/bin
/vendor
EOS

cat > .dockerignore << 'EOS'
EOS

git init
git add .
git commit -a -m 'initial import'

set -x
NAME=$APPNAME make 

bin/$APPNAME

