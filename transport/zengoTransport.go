package transport

import (
	"bingo/endpoint/server"
	"bingo/middleware"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

func InitSvc() http.Handler {
	svcendpoint := server.MakeSvcEndpoint()
	svcendpoint = middleware.MakeLogMiddleWare(svcendpoint)
	httpServer := httptransport.NewServer(
		svcendpoint,
		server.DecodeSvc,
		server.EncodeSvc,
	)
	return httpServer
}
