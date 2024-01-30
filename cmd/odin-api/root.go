package main

import (
	"fmt"

	"github.com/Manni-MinM/odin/internal/config"
	"github.com/Manni-MinM/odin/internal/database"
	"github.com/Manni-MinM/odin/internal/handler"
	"github.com/Manni-MinM/odin/internal/model"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
    srv := echo.New()

    srv.Use(middleware.Logger())
	srv.Use(middleware.Recover())

    conf, err := config.Load()
    if err != nil {
        srv.Logger.Fatal(err)
    }

    apiConf := conf.API

    db, err := database.RedisConn(apiConf.Redis)
    if err != nil {
        srv.Logger.Fatal(err)
    }

    serverHealthHandler := handler.NewHTTPServerHealthHandler(model.NewRedisServerHealthRepo(db))

    api := srv.Group("/api")

    api.POST("/server/", serverHealthHandler.Create)
    api.GET("/server/", serverHealthHandler.Get)
    api.GET("/server/all/", serverHealthHandler.GetAll)

    srvConf := apiConf.Server
    srvAddr := fmt.Sprintf(":%v", srvConf.Port)

    err = srv.Start(srvAddr)
    if err != nil {
        srv.Logger.Fatal(err)
    }
}
