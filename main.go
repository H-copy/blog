package main

import (

	"fmt"
	"log"
	"net/http"
	"time"

	// "github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
		
	"blog/internal/model"
	"blog/internal/routers"
	"blog/pkg/setting"
	"blog/pkg/logger"
	"blog/global"
)

func init(){
	
	if err := setupSetting(); err != nil{
		log.Fatalf("init config err:%v", err)
	}

	if err := setupDBEngine(); err != nil{
		log.Fatalf("init DB err:%v", err)
	}

	setupLogger()
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

func setupDBEngine() error{
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	return err
}

func setupLogger() error{
	
	writer := &lumberjack.Logger{
		Filename: global.AppSetting.LogSavePath + "/" +  global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize: 600,
		MaxAge: 10,
		LocalTime: true,
	}
	
	global.Logger = logger.NewLogger(writer, "", log.LstdFlags).WithCaller(2)
	
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

	global.Logger.Info("start server")
	
	
	s.ListenAndServe()
}