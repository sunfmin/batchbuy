echo "start run on dev server"

export GOROOT=/usr/local/go
export GOPATH=/usr/local/gopkg
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH

cd $GOPATH/src/github.com/sunfmin/batchbuy
git pull origin master
$GOROOT/bin/go install .

target="batchbuy"
echo 'kill running process'
killall $target;

cd /usr/local/gopkg/;
echo 'run in backgroud'
nohup $GOPATH/bin/$target >> /home/ubuntu/batchbuy.log 2>&1 &
