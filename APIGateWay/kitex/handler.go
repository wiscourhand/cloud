package main

import (
	"context"
	server0 "github.com/wiscourhand/rpcsvr/kitex_gen/kitex/test/server"
)

// ExampleServiceImpl implements the last service interface defined in the IDL.
type ExampleServiceImpl struct{}

// ExampleMethod implements the ExampleServiceImpl interface.
func (s *ExampleServiceImpl) ExampleMethod(ctx context.Context, req *server0.ExampleReq) (resp *server0.ExampleResp, err error) {
	// TODO: Your code here...
	resp = &server0.ExampleResp{
		Msg:      "call rpc2.0 success",
		BaseResp: nil,
	}
	return resp, nil
}
