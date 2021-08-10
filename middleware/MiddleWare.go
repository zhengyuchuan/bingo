package middleware

import (
	"bingo/component"
	"context"
	"github.com/go-kit/kit/endpoint"
)

func MakeLogMiddleWare(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		component.GetLogger().Info("日志中间件")
		return next(ctx, request)
	}
}
