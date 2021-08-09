package transport

import (
	"bingo/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)


func InitSvc() http.Handler {
	svcendpoint := endpoint.MakeSvcEndpoint()
	server := httptransport.NewServer(
		svcendpoint,
		endpoint.DecodeSvc,
		endpoint.EncodeSvc,
		)
	return server
}



