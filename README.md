# go-zero and vue project

## 1、创建服务

> goctl api new product-service

## 2、运行服务

> go run productservice.go -f etc/productservice-api.yaml
这里需要注意的是，需要到服务目录下执行

端口号修改在product-service下etc/productservice-api.yaml

生成代码:
> goctl api go --api product_service.api --dir ./

## 3、多个api文件产生冲突时

这里如果定义多个api文件，生成文件时会internal/handler里面的route文件会将其他的api文件覆盖掉
解决办法，定义一个index.api或者别的名字的文件，其中使用import引入其他的文件
生成时 只生成index.api文件
