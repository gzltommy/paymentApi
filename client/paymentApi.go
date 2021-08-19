package client

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gzltommy/common"
	"github.com/micro/go-micro/v2/client"
)

type clientWrapper struct {
	client.Client // 继承该接口
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	return hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
		// run 正常的执行逻辑
		common.Info(req.Service() + "." + req.Endpoint())
		return c.Client.Call(ctx, req, rsp, opts...)

	}, func(err error) error {
		common.Info(err)
		return err
	})
}

func NewClientHystrixWrapper() client.Wrapper {
	return func(c client.Client) client.Client {
		return &clientWrapper{Client: c}
	}
}
