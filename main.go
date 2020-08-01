package main

import (

	"fmt"
	"log"
	"net/http"
	"time"

	// "github.com/gin-gonic/gin"

	"blog/internal/routers"
	"blog/pkg/setting"
	"blog/global"
)

func init(){
	
	if err := setupSetting(); err != nil{
		log.Fatalf("init config err:%v", err)
	}
}

func setupSetting() error{
	setting, err := setting.NewSetting()
	if err != nil{
		return err
	}

	conf := map[string]interface{}{
		"Server": &global.ServerSetting,
		"App": &global.AppSetting,
		"Database": &global.DatabaseSetting,
	}

	for key, val := range conf{

		err := setting.ReadSection(key, val)

		if err != nil{
			return err
		}
		
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriterTimeout *= time.Second

	return nil
}

func main(){

	r := routers.NewRouter()
	s := &http.Server{
		Addr: fmt.Sprintf(":%d", global.ServerSetting.HttpPort),
		Handler: r,
		ReadTimeout: global.ServerSetting.ReadTimeout,
		WriteTimeout: global.ServerSetting.WriterTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}