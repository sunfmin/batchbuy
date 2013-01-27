Write API in api

# Installation


```sh
mkdir low_tea_at_the_plant low_tea_at_the_plant/bin low_tea_at_the_plant/pkg
cd low_tea_at_the_plant
git clone https://github.com/sunfmin/batchbuy
mv batchbuy src
export GOPATH=$(pwd):$GOPATH
export PATH=$(pwd):$PATH
go get github.com/gorilla/schema
go get labix.org/v2/mgo
go install controller
controller
```