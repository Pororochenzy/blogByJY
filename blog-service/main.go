package main

import (
	"github.com/go-programming-tour-book/blog-service/internal/routers"
	"net/http"
	"time"
)

func main(){
	router := routers.NewRouter()
	s:=&http.Server{
		Addr:":8080",
		Handler:router,
		ReadHeaderTimeout:10*time.Second,
		WriteTimeout:10*time.Second,
		MaxHeaderBytes:1<<20,
	}
	s.ListenAndServe()
	//r:=gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200,gin.H{"message":"pong"})
	//})
	//r.Run()
}
