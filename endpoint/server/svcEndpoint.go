package server

import (
	"bingo/endpoint/client"
	"bingo/service"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
	"net/http"
)

func EncodeSvc(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	//if f, ok := response.(httpendpoint.Failer); ok && f.Failed() != nil {
	//	errorEncoder(ctx, f.Failed(), w)
	//	return nil
	//}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func DecodeSvc(ctx context.Context, req *http.Request) (interface{}, error) {
	var request service.SvcRequest
	varMap := mux.Vars(req)
	request = service.SvcRequest{
		Name: varMap["name"],
		Age:  varMap["age"],
	}
	//err := json.NewDecoder(req.Body).Decode(&request)
	//if err != nil {
	//	return nil, err
	//} else {
	//	return request, nil
	//}
	return request, nil
}

func MakeSvcEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		svcrequest, ok := request.(service.SvcRequest)
		if !ok {
			return nil, errors.New("解析失败！")
		}
		juejinEndpoint := client.MakeJuejinEndpoint()
		svcresponse := service.SvcService(ctx, svcrequest, juejinEndpoint)
		return svcresponse, nil
	}
}
