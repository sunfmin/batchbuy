echo "start run on dev server"

export GOROOT=/opt/go/bin/go
export GOPATH=/home/lowtea/goprojs
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH

cd $GOPATH/src/github.com/sunfmin/batchbuy
git pull origin master
$GOROOT/bin/go install .

target="batchbuy"
echo 'kill running process'
killall $target;

cd $GOPATH;
echo 'run in backgroud'
nohup $GOPATH/bin/$target >> /home/lowtea/goprojs/batchbuy.log 2>&1 &
