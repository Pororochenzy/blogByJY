package main

import (
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/routers"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSeeting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	//fmt.Println(global.ServerSetting.)
}
func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:              ":8080",
		Handler:           router,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
	s.ListenAndServe()
	//r:=gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200,gin.H{"message":"pong"})
	//})
	//r.Run()
}

func setupSeeting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}
