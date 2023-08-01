package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type IdlReq struct {
	idlName string
	idlInfo IdlInfo
}

type IdlInfo struct {
	content  string
	includes map[string]string
}

/*
添加idl信息，用于测试idl热更新
*/
func main() {
	url := flag.String("url", "http://127.0.0.1:8888/idl-info/add/CServiceName", "request url")

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

	var idlInfo = IdlInfo{
		content:  content,
		includes: includes,
	}

	var idlReq = IdlReq{
		idlName: "CServiceName",
		idlInfo: idlInfo,
	}

	dataType, _ := json.Marshal(idlReq)
	idlReqStr := string(dataType)

	reqBody := flag.String("reqBody", idlReqStr, "json data")

	flag.Parse()

	post(*url, "application/json", *reqBody)

}

/*
发送post请求
*/
func post(url string, contentType string, reqBody string) {
	fmt.Println("POST REQ...")
	fmt.Println("REQ:", reqBody)
	client := http.Client{}
	rsp, err := client.Post(url, contentType, strings.NewReader(reqBody))
	if err != nil {
		fmt.Println(err)
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("RSP:", string(body))
}

func (info IdlInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"content":  info.content,
		"includes": info.includes,
	})
}

func (req IdlReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"idlName": req.idlName,
		"idlInfo": req.idlInfo,
	})
}
