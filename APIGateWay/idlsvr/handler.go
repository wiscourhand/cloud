package main

import (
	"context"
	demo "github.com/wiscourhand/idlsvr/kitex_gen/demo"
)

// IdlServiceImpl implements the last service interface defined in the IDL.
type IdlServiceImpl struct{}

// 用来存储idl信息的图
var idlInfoMap = make(map[string]*demo.IdlInfo)

// Register implements the IdlServiceImpl interface.
func (s *IdlServiceImpl) Register(ctx context.Context, idlReq *demo.IdlReq) (resp *demo.AddIdlResp, err error) {
	// TODO: Your code here...

	//将新的idl信息储存起来
	idlInfoMap[idlReq.IdlName] = idlReq.IdlInfo

	resp = new(demo.AddIdlResp)

	//返回添加成功的信息
	resp.Message = "add idl success " + idlReq.IdlName
	resp.Success = true

	return resp, nil
}

// Query implements the IdlServiceImpl interface.
func (s *IdlServiceImpl) Query(ctx context.Context, name string) (resp *demo.IdlInfo, err error) {
	// TODO: Your code here...

	resp = new(demo.IdlInfo)

	//从储存中获取idl信息
	resp = idlInfoMap[name]

	return resp, nil
}
