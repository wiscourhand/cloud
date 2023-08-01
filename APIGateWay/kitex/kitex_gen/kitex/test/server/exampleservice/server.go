// Code generated by Kitex v0.6.1. DO NOT EDIT.
package exampleservice

import (
	server "github.com/cloudwego/kitex/server"
	server0 "github.com/wiscourhand/rpcsvr/kitex_gen/kitex/test/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler server0.ExampleService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
