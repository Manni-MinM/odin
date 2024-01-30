package healthcheck

import (
    "os"
    "log"
    "time"
	"net/http"

	"github.com/Manni-MinM/odin/internal/config"
	"github.com/Manni-MinM/odin/internal/model"
)

type HealthChecker interface {
    Healthy(string) (bool, error)
    Checkup()
}

type HTTPHealthCheck struct {
    client      *http.Client
    repo        model.ServerHealthRepo
    logger      *log.Logger
}

func NewHTTPHealthCheck(repo model.ServerHealthRepo, conf config.Cron) *HTTPHealthCheck {
    client := &http.Client{
        Timeout: time.Duration(conf.Timeout) * time.Second,
    }

    logger := log.New(os.Stdout, "[HealthChecker] ", log.Ldate|log.Ltime)

    return &HTTPHealthCheck{client, repo, logger}
}

func (hc *HTTPHealthCheck) Healthy(url string) (bool, error) {
    resp, err := hc.client.Get(url)
    if err != nil {
        return false, err
    }

    if resp.StatusCode != http.StatusOK {
        return false, nil
    }

    return true, nil
}

func (hc *HTTPHealthCheck) Checkup() {
    shMap, err := hc.repo.GetAll()
    if err != nil {
        hc.logger.Println(err)
    } 

    for id, sh := range(shMap) {
        isHealthy, err := hc.Healthy(sh.Address)
        if err != nil {
            hc.logger.Println(err)
        } else if isHealthy {
            hc.logger.Printf("Get request on \"%v\" successful\n", sh.Address)
        } else {
            hc.logger.Printf("Get request on \"%v\" failed\n", sh.Address)
        }

        switch isHealthy {
        case true:
            err = hc.repo.IncrementSuccessByID(id)
        case false:
            err = hc.repo.IncrementFailureByID(id)
        }

        if err != nil {
            hc.logger.Println(err)
        }
    }
}
