package transport

import (
	"bingo/endpoint"
	"bingo/middleware"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

func InitSvc() http.Handler {
	svcendpoint := endpoint.MakeSvcEndpoint()
	svcendpoint = middleware.MakeLogMiddleWare(svcendpoint)
	server := httptransport.NewServer(
		svcendpoint,
		endpoint.DecodeSvc,
		endpoint.EncodeSvc,
	)
	return server
}
