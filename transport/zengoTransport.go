package transport

import (
	"bingo/endpoint/server"
	"bingo/middleware"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

type ServiceSet struct {
	SvcService http.Handler
	StuService http.Handler
}

func InitSvc() ServiceSet {
	var ServiceIpl ServiceSet
	{
		svcendpoint := server.MakeSvcEndpoint()
		svcendpoint = middleware.MakeLogMiddleWare(svcendpoint)
		svcServer := httptransport.NewServer(
			svcendpoint,
			server.DecodeSvc,
			server.EncodeSvc,
		)
		ServiceIpl.SvcService = svcServer
	}

	return ServiceIpl
}
