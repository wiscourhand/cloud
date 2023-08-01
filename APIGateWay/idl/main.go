package main

import (
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/wiscourhand/idl/kitex_gen/kitex/test/server/exampleservice"
	"log"
	"net"

	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	//设置监听端口
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8891")

	//注册服务
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	//设置注册信息
	ri := &registry.Info{
		ServiceName: "CServiceName",
		Tags: map[string]string{
			"Cluster": "xxx",
		},
	}

	//构建服务器
	svr := exampleservice.NewServer(
		new(ExampleServiceImpl),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "CServiceName",
		}),
		server.WithRegistryInfo(ri),
		server.WithServiceAddr(addr),
	)

	//启动服务
	err = svr.Run()
	if err != nil {
		panic(err)
	}
}
