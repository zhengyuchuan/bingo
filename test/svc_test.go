package test

import (
	"bingo/transport"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSvc(t *testing.T) {

	handler := transport.InitSvc()
	server := httptest.NewServer(handler)
	defer server.Close()

	reader := strings.NewReader(`{"name": "zheng", "age": 15}`)
	resp, _ := http.NewRequest(http.MethodGet, server.URL, reader)

	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, resp)

	respJson, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		fmt.Println("解析失败！")
	}

	fmt.Println(string(respJson))

}
