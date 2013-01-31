Write API in api

# Installation


```sh
export GOPATH=path/to/your/go/workspace
cd path/to/your/go/workspace
go get -u github.com/sunfmin/batchbuy
go install github.com/sunfmin/batchbuy
mongoimport --db low_tea_at_the_plant --collection products --file /src/github.com/sunfmin/batchbuy/dump/products.json
bin/batchbuy
```