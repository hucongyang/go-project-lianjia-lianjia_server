package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/global"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/internal/model"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/internal/routers"
	"github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/logger"
	setting2 "github.com/hucongyang/go-project-lianjia-lianjia_server/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"syscall"
	"time"
)

func init() {
	// 配置信息初始化
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	// 数据库连接初始化
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	// 日志信息初始化
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

// @title 链家二手房房价数据api接口服务系统
// @version 1.0
// @description
// @termsOfService https://github.com/hucongyang/go-project-lianjia-lianjia_server
func main() {
	gin.SetMode(global.ServerSetting.RunMode)

	router := routers.NewRouter()
	readTimeout := global.ServerSetting.ReadTimeout
	writeTimeout := global.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", global.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20
	//server := &http.Server{
	//	Addr:           endPoint,
	//	Handler:        router,
	//	ReadTimeout:    readTimeout,
	//	WriteTimeout:   writeTimeout,
	//	MaxHeaderBytes: maxHeaderBytes,
	//}
	//log.Printf("[info] start http server listening %s", endPoint)
	//server.ListenAndServe()

	// 2. 优雅的重启 使用 endless
	endless.DefaultReadTimeOut = readTimeout
	endless.DefaultWriteTimeOut = writeTimeout
	endless.DefaultMaxHeaderBytes = maxHeaderBytes
	server := endless.NewServer(endPoint, router)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server error: %v", err)
	}

	//// 1. 优雅的停止：更新服务不停止现有服务
	//go func() {
	//	if err := service.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	//		log.Fatalf("service.ListenAndServe error: %v", err)
	//	}
	//}()
	//// 等待中断信号
	//quit := make(chan os.Signal)
	//// 接收 syscall.SIGINT 和 syscall.SIGTERM 信号
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//<-quit
	//log.Println("Shuting down server...")
	//// 最大时间控制，通知该服务端它有5s的时间来处理原有的请求
	//// 如果没有正在处理的旧请求，那么在按组合键ctrl+c后，其会直接退出（因为不需要等待）
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//if err := service.Shutdown(ctx); err != nil {
	//	log.Fatal("Server forced to shutdown: ", err)
	//}
	//log.Println("Server exiting")
}

// 初始化配置
func setupSetting() error {
	setting, err := setting2.NewSetting()
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

// 初始化数据库连接
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

// 初始化日志
func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   500,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
