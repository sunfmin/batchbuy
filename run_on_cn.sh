#!/usr/bin/env bash

set -e

echo "start run on dev server"

export GOROOT=/opt/go
export GOPATH=/home/lowtea/goprojs
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH

# cd $GOPATH/src/github.com/sunfmin/batchbuy
# git pull origin master
# sudo GOPATH=/home/lowtea/goprojs $GOROOT/bin/go install -a

target="batchbuy"
echo 'kill running process'
killall $target;

cd /home/lowtea
tar mxf batchbuy_linux_amd64.tar.gz
cp batchbuy_linux_amd64/batchbuy $GOPATH/bin/$target

cd $GOPATH;
echo 'run in backgroud'
nohup $GOPATH/bin/$target >> /home/lowtea/goprojs/batchbuy.log 2>&1 &

tail -f /home/lowtea/goprojs/batchbuy.log
