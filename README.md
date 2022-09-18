# Kratos Project Template

## 项目说明

kratos-news-system通过实现用户模块，新闻文章模块，评论模块去探索kratos的基本设计理念，包括项目的初始化、kratos工具的使用，makefile文件的编写提升操作的便捷性。API接口的定义，配置的使用、HTTP和GRPC的使用。

## 参考资料
```api
grpc UnimplementedXxxServer妙用，通过UnimplementedXxxServer去实现定义的rpc服务，实例化UnimplementedXxxServer，通过嵌入组合的方式引入UnimplementedXxxServer，实现要完成的功能函数覆盖之前默认生成的。
https://blog.csdn.net/qq_17199495/article/details/125910932

kratos wire依赖注入详解，了解kratos通过wire生成wire_gen.go文件，将要依赖注入的方法，统一注入
https://go-kratos.dev/blog/go-project-wire

采用gorm实现mysql的实例及操作
https://gorm.io/zh_CN/docs/

redis的基本使用
https://redis.uptrace.dev/guide/go-redis.html
```

## Install Kratos
```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```
## Create a service
```
# Create a template project
kratos new server

cd server
# Add a proto template
kratos proto add api/server/server.proto
# Generate the proto code
kratos proto client api/server/server.proto
# Generate the source code of service by proto file
kratos proto server api/server/server.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs
```
## Generate other auxiliary files by Makefile
```
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```
## Automated Initialization (wire)
```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```

## Docker
```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```

