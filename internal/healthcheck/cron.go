package healthcheck

import (
	"context"

	"github.com/Manni-MinM/odin/internal/config"

	"github.com/robfig/cron/v3"
)

type HealthCheckScheduler struct {
    pattern         string
    scheduler       *cron.Cron
    checker         HealthChecker
}

func NewHealthCheckScheduler(checker HealthChecker, conf config.Cron) *HealthCheckScheduler {
    return &HealthCheckScheduler {
        pattern: conf.Pattern,
        scheduler: cron.New(cron.WithSeconds()),
        checker: checker,
    }
}

func (hcs *HealthCheckScheduler) Start() error {
    _, err := hcs.scheduler.AddFunc(hcs.pattern, hcs.checker.Checkup)
    if err != nil {
        return err
    }

    hcs.scheduler.Start()

    return nil
}

func (hcs *HealthCheckScheduler) Stop() context.Context {
    return hcs.scheduler.Stop()
}
