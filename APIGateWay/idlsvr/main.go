package main

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/wiscourhand/idlsvr/kitex_gen/demo"
	"github.com/wiscourhand/idlsvr/kitex_gen/demo/idlservice"

	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")

	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	idlSvr := new(IdlServiceImpl)

	svr := idlservice.NewServer(
		idlSvr,
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "idlSvr",
		}),
		server.WithServiceAddr(addr),
	)

	//初始化储存idl的图
	initIdlMap()

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

/*
提前准备部分idl信息便于测试
*/
func initIdlMap() {

	path := "a/b/main.thrift"
	content := `
include "base.thrift"
namespace go kitex.test.server

struct ExampleReq {
    1: required string Msg,
    255: base.Base Base,
}
struct ExampleResp {
    1: required string Msg,
    255: base.BaseResp BaseResp,
}
service ExampleService {
    ExampleResp ExampleMethod(1: ExampleReq req)(api.post = '/add-student-info'),
}

	`
	includes := map[string]string{
		path: content,
		"base.thrift": `
namespace py base
namespace go base
namespace java com.xxx.thrift.base

struct TrafficEnv {
    1: bool Open = false,
    2: string Env = "",
}

struct Base {
    1: string LogID = "",
    2: string Caller = "",
    3: string Addr = "",
    4: string Client = "",
    5: optional TrafficEnv TrafficEnv,
    6: optional map<string, string> Extra,
}

struct BaseResp {
    1: string StatusMessage = "",
    2: i32 StatusCode = 0,
    3: optional map<string, string> Extra,
}

`,
	}

	var svrInfo = new(demo.IdlInfo)

	svrInfo.Content = content
	svrInfo.Includes = includes

	idlInfoMap["AServiceName"] = svrInfo
	idlInfoMap["BServiceName"] = svrInfo
}
