package main

import (
    "os"
    "log"
    "syscall"
    "os/signal"

	"github.com/Manni-MinM/odin/internal/model"
	"github.com/Manni-MinM/odin/internal/config"
	"github.com/Manni-MinM/odin/internal/database"
	"github.com/Manni-MinM/odin/internal/healthcheck"
)

func main() {
    logger := log.New(os.Stdout, "[Healthcheck] ", log.Ldate|log.Ltime)

    logger.Println("Initiaing Healthcheck service")

    conf, err := config.Load()
    if err != nil {
        logger.Fatal(err)
    }

    hcConf := conf.HealthCheck

    db, err := database.RedisConn(hcConf.Redis)
    if err != nil {
        logger.Fatal(err)
    }

    repo := model.NewRedisServerHealthRepo(db)

    scheduler := healthcheck.NewHealthCheckScheduler(
        healthcheck.NewHTTPHealthCheck(repo, hcConf.Cron),
        hcConf.Cron,
    )

    go func() {
        err := scheduler.Start()
        if err != nil {
            logger.Fatal(err)
        }
    }()

    signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

    signal := <-signalCh
    logger.Printf("received shutdown signal: %v. initiating graceful shutdown\n", signal)

    scheduler.Stop()
}
