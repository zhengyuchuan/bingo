package client

import (
	"bingo/service"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	httpendpoint "github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"io/ioutil"
	"net/http"
	"net/url"
)

func EncodeJuejin(cxt context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Cookie", "_ga=GA1.2.2010206452.1620609324; n_mh=g5uT9tzMjksycE0DN2xLd5pvdFboWpgOLRQtLQiy8Ds; passport_csrf_token_default=3dd15882c7ce7c6882c053534c2b3cdc; passport_csrf_token=3dd15882c7ce7c6882c053534c2b3cdc; sid_guard=609b8fb1965fb354caa2977cc6305229%7C1625794051%7C5184000%7CTue%2C+07-Sep-2021+01%3A27%3A31+GMT; uid_tt=d8517cfa25186846d36ec14145f9ce44; uid_tt_ss=d8517cfa25186846d36ec14145f9ce44; sid_tt=609b8fb1965fb354caa2977cc6305229; sessionid=609b8fb1965fb354caa2977cc6305229; sessionid_ss=609b8fb1965fb354caa2977cc6305229; MONITOR_WEB_ID=8634b3b5-ead5-47cc-b54c-fde4ed8d1f90; _gid=GA1.2.1348569692.1628472302")
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func DecodeJuejin(ctx context.Context, r *http.Response) (response interface{}, err error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp service.JuejinResponse
	err = json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

func MakeJuejinEndpoint() httpendpoint.Endpoint {
	return httptransport.NewClient(
		"POST",
		getJuejinUrl(),
		EncodeJuejin,
		DecodeJuejin,
	).Endpoint()
}

func getJuejinUrl() *url.URL {
	juejinUrl, err := url.Parse("https://api.juejin.cn/recommend_api/v1/short_msg/recommend")
	if err != nil {
		return nil
	}
	return juejinUrl
}
