package main

import (
	"bingo/transport"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	flag "github.com/ogier/pflag"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var env = flag.StringP("config", "c", "test", "environment")
var configPathMap = map[string]string{"test": "conf/test.yaml", "prod": "conf/prod.yaml"}

func main() {
	flag.Parse()
	configPath := configPathMap[*env]
	if configPath == "" {
		configPath = configPathMap["test"]
	}

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
	serverUrl := viper.GetString("server_url")
	fmt.Println(serverUrl)

	// TODO：初始化各种组件

	// TODO：初始化路由

	// TODO: 初始化HandleFunc及middlewire
	r := mux.NewRouter()
	svcService := transport.InitSvc()
	r.Handle("/hello", svcService).Methods("GET")
	srv := http.Server{
		Addr:    ":8081",
		Handler: r,
	}
	exit := make(chan os.Signal)
	// 监听ctrl+c
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-exit
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := srv.Shutdown(ctx)
		if err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println(srv.ListenAndServe())
}
