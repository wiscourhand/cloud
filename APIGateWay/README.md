## 说明文档

## 1结构
├─.idea
├─httpsvr
│  ├─biz
│  │  ├─handler
│  │  └─router
│  ├─kitex_gen
│  │  └─demo
│  │      └─idlservice
│  └─script
├─idl
│  ├─.idea
│  ├─add_idl
│  ├─kitex_gen
│  │  ├─base
│  │  └─kitex
│  │      └─test
│  │          └─server
│  │              └─exampleservice
│  ├─output
│  │  └─bin
│  └─script
├─idlsvr
│  ├─.idea
│  ├─kitex_gen
│  │  └─demo
│  │      └─idlservice
│  ├─output
│  │  └─bin
│  └─script
├─kitex
│  ├─.idea
│  ├─kitex_gen
│  │  ├─base
│  │  └─kitex
│  │      └─test
│  │          └─server
│  │              └─exampleservice
│  ├─output
│  │  └─bin
│  └─script
└─rpcsvr
    ├─.idea
    ├─kitex_gen
    │  ├─base
    │  └─kitex
    │      └─test
    │          └─server
    │              └─exampleservice
    ├─output
    │  └─bin
    └─script

## 2接口

1  用于提供网关服务的接口

r.POST("/:svrName/:methodName", handler.RpcMethod)

func RpcMethod(ctx context.Context, c *app.RequestContext)

2  用于管理IDL资源的接口

r.POST("/idl-info/add/:idlName", handler.AddIdlInfo)

func AddIdlInfo(ctx context.Context, c *app.RequestContext)

## 2技术要求

Hertz:一个高可用性、高性能、高可扩展性的 HTTP 框架，支持微服务的开发。Hertz 提供了允许 API 网关解释和响应来自 客户端的 HTTP 请求的基础设施。

Kitex:一个高性能和强可扩展性的 Golang RPC框架，支持构建微服务。Kitex 为网关和提供服务的服务器之间的 RPC通信提供 了许多基础设施和特性。在我们的项目中，它允许我们创建多个微服务进行测试。

etcd:它是一个开源的、分布式的、一致的键值存储，用于共享配置、服务发现和分布式系统或机器集群的调度器协调。 在我们的项目中，我们使用它来注册多个基于 kitex 的服务，并使它们能够被 API 网关发现。

ApacheThrift:它是一种接口定义语言和二进制通信协议，用于为众多编程语言定义和创建服务。在这个项目中，我们创建了 thrift 文件，这些文件可以自动为 Hertz 和 Kitex 服务器生成脚手架代码，并将 RPC请求转换为 thrift 二进制格式。

Go: Go 是一种强大而高效的编程语言，它优先考虑简单性和并发性。我们选择 Go 作为项目的基础，利用它的优势来构建我们 的 API 网关。此外，我们还利用了专门为 Go 设计的 Hertz 和 Kitex 框架。

## 3部署步骤

1  开启etcd服务

2  开启httpSvr服务

3  开启idlSvr服务

4  开启各个rpcSvr服务

## 4测试

1 方案：基于Golang性能测试框架Benchmark及第三方性能测试工具Apache Benchmark进行

2  基准测试
   使用Benchmark向AServiceName发送请求，显示数据：
     共使用：4个核心
     处理次数：796次
     每次耗时：2355843ns

   并行测试
   基于Benchmark设置最大并发数为4时：
     共使用：4个核心
     处理次数：1359次
     每次耗时：1196297ns
   基于Benchmark设置最大并发数为8时：
     共使用：4个核心
     处理次数：15611次
     每次耗时：792671ns
   结论：并发数增加，处理次数增加，耗时减少

