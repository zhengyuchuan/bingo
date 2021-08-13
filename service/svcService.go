package service

import (
	"bingo/endpoint/client"
	"context"
)

type SvcRequest struct {
	Name string `http:"name" json:"name"`
	Age  string `http:"age" json:"age"`
}

type SvcResponse struct {
	Id   int           `json:"id"`
	Name string        `json:"name"`
	Age  int           `json:"age"`
	Data []client.Data `json:"data"`
}

func SvcService(ctx context.Context, svc SvcRequest) SvcResponse {
	juejinEndpoint := client.MakeJuejinEndpoint()
	data, err := juejinEndpoint(ctx, client.JuejinRequest{
		Cursor:   "0",
		IdType:   4,
		Limit:    20,
		SortType: 300,
	})
	if err != nil {
		return SvcResponse{
			Id:   0,
			Name: "zheng",
			Age:  10,
			Data: nil,
		}
	}
	newData, ok := data.(client.JuejinResponse)
	if !ok {
		return SvcResponse{
			Id:   0,
			Name: "zheng",
			Age:  10,
			Data: nil,
		}
	}

	return SvcResponse{
		Id:   0,
		Name: "zheng",
		Age:  10,
		Data: newData.Data,
	}
}
