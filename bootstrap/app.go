package bootstrap

import (
	"context"
	"github.com/gin-gonic/gin"
	"github/siafei/gin-test/global"
	"github/siafei/gin-test/pkg/databases"
	"github/siafei/gin-test/pkg/logger"
	"github/siafei/gin-test/pkg/redis"
	"github/siafei/gin-test/pkg/setting"
	"github/siafei/gin-test/route"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type app struct {
}

func NewApp() *app {
	if err := InitSetting(); err != nil {
		panic(err.Error())
	}
	if err := InitDb(); err != nil {
		panic(err.Error())
	}
	if err :=setupLogger(); err != nil {
		panic(err.Error())
	}
	InitRedis()
	return &app{}
}

/**
数据库初始化
*/
func InitDb() error {
	var err error
	if global.DB, err = databases.NewDb(global.DatabaseSetting); err != nil {
		return err
	}
	if global.Db_test2, err = databases.NewDb(global.DatabaseSetting2); err != nil {
		return err
	}
	return nil
}

/**
初始化redis
 */
func InitRedis() {
	global.Redis  = redis.NewRedis(global.RedisSetting)
}



/**
错误日志初始化
*/
func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

/**
配置初始化
*/
func InitSetting() error {
	set, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = set.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = set.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = set.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = set.ReadSection("Database2", &global.DatabaseSetting2)
	if err != nil {
		return err
	}
	err = set.ReadSection("Redis", &global.RedisSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func (a app) Run() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := route.NewRoute()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		WriteTimeout:   10 * time.Second,
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("s.ListenAndServe err: %v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Start Shut Down Server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if error := s.Shutdown(ctx); error != nil {
		log.Fatal("Server forced to shutdown:", error)
	}
	log.Fatal("Server Shut Down!")
}
