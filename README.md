# Kratos News System

## 项目说明

kratos-news-system通过实现用户模块，新闻文章模块，评论模块去探索kratos的基本设计理念，包括项目的初始化、kratos工具的使用，makefile文件的编写提升操作的便捷性。API接口的定义，配置的使用、HTTP和GRPC的使用,使用consul服务发现、jaeger链路追踪、kafka消息队列。

## 参考资料
```api
grpc UnimplementedXxxServer妙用，通过UnimplementedXxxServer去实现定义的rpc服务，实例化UnimplementedXxxServer，通过嵌入组合的方式引入UnimplementedXxxServer，实现要完成的功能函数覆盖之前默认生成的。
https://blog.csdn.net/qq_17199495/article/details/125910932

kratos wire依赖注入详解，了解kratos通过wire生成wire_gen.go文件，将要依赖注入的方法，统一注入
https://go-kratos.dev/blog/go-project-wire

kratos启动流程明细
https://go-kratos.dev/blog/go-layout-operation-process

采用gorm实现mysql的实例及操作
https://gorm.io/zh_CN/docs/

redis的基本使用
https://redis.uptrace.dev/guide/go-redis.html
```
