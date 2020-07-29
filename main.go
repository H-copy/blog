package main

import (

	// "fmt"
	"net/http"
	"time"

	// "github.com/gin-gonic/gin"

	"blog/internal/routers"
	
)

func main(){

	r := routers.NewRouter()
	s := &http.Server{
		Addr: ":3000",
		Handler: r,
		ReadTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}