// Code generated by Kitex v0.6.1. DO NOT EDIT.

package exampleservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	server0 "github.com/wiscourhand/idl/kitex_gen/kitex/test/server"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	ExampleMethod(ctx context.Context, req *server0.ExampleReq, callOptions ...callopt.Option) (r *server0.ExampleResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kExampleServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kExampleServiceClient struct {
	*kClient
}

func (p *kExampleServiceClient) ExampleMethod(ctx context.Context, req *server0.ExampleReq, callOptions ...callopt.Option) (r *server0.ExampleResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ExampleMethod(ctx, req)
}
