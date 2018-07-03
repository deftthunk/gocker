#!/bin/bash

# NOTE: build command will place built binaries in current working dir
# if you are in/below your GOPATH. otherwise it will place them in GOPATH/GOBIN
# as is normally expected


#----------------CHANGEME---------------------CHANGEME-------------------------
# host's $HOME relative GOPATH (ie. /home/user/code/go -> 'code/go')
godir='code/go'

# use Alpine Linux (~400MB) or Debian aka 'latest' (~800MB) as base
# they should behave identically except in the case of using CGO, where Alpine 
# will use MUSL libc instead of glibc. This only matters when using libraries 
# that require dynamic linking. Both Alpine and Debian will point to the latest
# stable Go version.
#
# TLDR: If you're not including C code, leave as 'alpine'. Else, put 'latest'
version='latest'
#----------------CHANGEME---------------------CHANGEME-------------------------


# what base are we using? download if not present
if [[ $version -eq 'alpine' ]]; then
	if [[ ! `docker image ls -q golang:alpine` ]]; then
		docker pull golang:alpine
	fi
elif [[ ! `docker image ls -q golang:latest` ]]; then
		docker pull golang:latest
fi

# default container working dir path
workdir='/go/bin'

# if current dir is under $godir, reassign container working dir
# to mirror host working dir path
if [[ $PWD =~ $godir(.*) ]]; then
	workdir="/go${BASH_REMATCH[1]}"
fi

# quick version that doesn't mount volumes
gohelp="docker run --rm --label 'golang' -u 1000 \
        golang:$version go help"

# mount arguments for docker volumes
gomnt=" --mount type=bind,source=$HOME/$godir,target=/go"

# docker environment vars
goenv=" -e GOBIN=/go/bin \
        -e GOTMPDIR=/tmp \
        -e GOCACHE=/tmp/.cache/go-build"

# full Go command; 
gocmd="docker run --rm --label golang -u 1000 \
       -w $workdir $gomnt $goenv --tmpfs /tmp:exec,mode=777,size=65535k golang:$version go"

# if container runs without args or has "help" as a CLI arg,
# run 'go'. else, run the full thing with volumes mounted
if [[ $# -eq 0 || $1 == 'help' ]]; then
	$gohelp
elif [[ "$#" -ge 1 ]]; then
	$gocmd $@
fi
