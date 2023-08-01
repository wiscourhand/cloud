package handler

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/adaptor"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	kclient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/wiscourhand/httpsvr/kitex_gen/demo/idlservice"
)

var cliMap = make(map[string]genericclient.Client)

var idlSvr genericclient.Client

func RpcMethod(ctx context.Context, c *app.RequestContext) {
	httpReq, err := adaptor.GetCompatRequest(c.GetRequest())
	if err != nil {
		panic("get http req failed")
	}

	customReq, err := generic.FromHTTPRequest(httpReq)
	if err != nil {
		panic(err)
		//panic("get custom req failed")
	}

	//转json
	jsonStr, err := json.Marshal(customReq.Body)
	jsonReq := string(jsonStr)

	//从路由中获取目标服务名
	targetName := c.Param("svrName")

	//在缓存中查找相关服务的客户端
	cli := cliMap[targetName]

	//如果缓存中不存在就向idlSvr请求
	if cli == nil {
		cli = getIdlInfo(ctx, targetName)
	}

	//如果idlSvr也中不存在就返回
	if cli == nil {
		c.JSON(consts.StatusOK, "can not find "+targetName)
		return
	}

	//发送泛化调用请求
	resp, err := cli.GenericCall(ctx, c.Param("methodName"), jsonReq)

	c.JSON(consts.StatusOK, resp)
}

func AddIdlInfo(ctx context.Context, c *app.RequestContext) {
	httpReq, err := adaptor.GetCompatRequest(c.GetRequest())
	if err != nil {
		panic("get http req failed")
	}

	customReq, err := generic.FromHTTPRequest(httpReq)
	if err != nil {
		panic(err)
		//panic("get custom req failed")
	}

	//转json
	jsonStr, err := json.Marshal(customReq.Body)
	jsonReq := string(jsonStr)

	//构建idlSvr的客户端
	if idlSvr == nil {
		idlSvr = initIdlGenericClient("idlSvr")
	}

	cli := idlSvr

	//发送泛化调用请求
	resp, err := cli.GenericCall(ctx, "Register", jsonReq)

	//向idlSvr获取idl信息
	getIdlInfo(ctx, c.Param("idlName"))

	c.JSON(consts.StatusOK, resp)
}

/*
向idlSvr获取idl信息
*/
func getIdlInfo(ctx context.Context, name string) genericclient.Client {

	//构建idlSvr客户端
	idlCli, err := idlservice.NewClient("idlSvr", kclient.WithHostPorts("127.0.0.1:9999"))
	if err != nil {
		panic("err new client:" + err.Error())
	}

	//发送查询请求
	idlResp, err := idlCli.Query(ctx, name)

	//不存在该服务的idl信息
	if idlResp == nil {
		return nil
	}

	if err != nil {
		panic("err query rpc server:" + err.Error())
	}

	//构建该服务的泛化调用客户端
	newCli := initGenericClient(name, idlResp.Content, idlResp.Includes)

	//将新生成的泛化调用客户端存入缓存
	cliMap[name] = newCli

	return cliMap[name]
}

/*
构建泛化调用客户端
*/
func initGenericClient(svrName string, content string, includes map[string]string) genericclient.Client {

	//根据内容解析IDL文件
	p, err := generic.NewThriftContentProvider(content, includes)
	if err != nil {
		panic(err)
	}

	// 构造 json 类型的泛化调用
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}

	//从注册中心获取目标服务
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic("err in resolver")
		//log.Fatal(err)
	}

	//构建泛化调用客户端，采用加权随机负载均衡
	cli, err := genericclient.NewClient(
		svrName,
		g,
		kclient.WithResolver(r),
		kclient.WithTag("Cluster", "xxx"),
		kclient.WithLoadBalancer(loadbalance.NewWeightedRandomBalancer()),
	)
	if err != nil {
		//return nil
		panic(err)
	}
	return cli
}

/*
构建idlSvr的客户端
*/
func initIdlGenericClient(svrName string) genericclient.Client {

	//idlSvr的idl文件
	path := "a/b/main.thrift"
	content := `
namespace go demo



struct IdlInfo {
    1: string content(go.tag = 'json:"content"'),
    2: map<string, string> includes(go.tag = 'json:"includes"'),
}

struct IdlReq {
    1: string idlName(api.body='name'),
    2: IdlInfo idlInfo(api.body='info'),
}

struct AddIdlResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

service IdlService {
    AddIdlResp Register(1: IdlReq idlReq)(api.post = '/add-ldl-info')
    IdlInfo Query(1: string name)(api.get = '/query')
}
	`
	includes := map[string]string{
		path: content,
	}

	//根据内容解析IDL文件
	p, err := generic.NewThriftContentProvider(content, includes)
	if err != nil {
		panic(err)
	}

	// 构造 json 类型的泛化调用
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}

	//从注册中心获取idlSvr服务
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic("err in resolver")
		//log.Fatal(err)
	}

	cli, err := genericclient.NewClient(
		svrName,
		g,
		kclient.WithResolver(r),
		kclient.WithTag("Cluster", "xxx"),
		kclient.WithLoadBalancer(loadbalance.NewWeightedRandomBalancer()),
	)
	if err != nil {
		panic(err)
	}
	return cli
}
