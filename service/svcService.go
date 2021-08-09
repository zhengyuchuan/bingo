package service

type SvcRequest struct {
	Name string `http:"name" json:"name"`
	Age string `http:"age" json:"age"`
}

type SvcResponse struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}

func SvcService(svc SvcRequest) SvcResponse {
	return SvcResponse{
		Id: 0,
		Name: "zheng",
		Age: 10,
	}
}
