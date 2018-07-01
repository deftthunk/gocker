#!/bin/bash

# NOTE: build command will place built binaries in current working dir
# if you are in your GOPATH. otherwise it will place them in GOPATH/GOBIN
# as is normally expected

# host's relative GOPATH (CHANGEME)
godir='code/go'
# default container working dir path
workdir='/go/bin'

# if current dir is under $godir, reassign container working dir
# to mirror host working dir path
if [[ $PWD =~ $godir(.*) ]]; then
	workdir="/go${BASH_REMATCH[1]}"
fi

# quick version that doesn't mount volumes
gohelp="docker run --rm --label 'golang' -u 1000 \
        golang:alpine go help"

# mount arguments for docker volumes
gomnt=" --mount type=bind,source=$HOME/$godir,target=/go"

# docker environment vars
goenv=" -e GOBIN=/go/bin \
        -e GOTMPDIR=/tmp \
        -e GOCACHE=/tmp/.cache/go-build"

# full Go command; 
gocmd="docker run --rm --label golang -u 1000 \
       -w $workdir $gomnt $goenv --tmpfs /tmp:exec,mode=777,size=65535k golang:alpine go"

# if container runs without args or has "help" as a CLI arg,
# run 'go'. else, run the full thing with volumes mounted
if [[ $# -eq 0 || $1 == 'help' ]]; then
	$gohelp
elif [[ "$#" -ge 1 ]]; then
	$gocmd $@
fi
