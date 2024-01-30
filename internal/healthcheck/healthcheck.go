package healthcheck

import (
    "time"
	"net/http"

	"github.com/Manni-MinM/odin/internal/config"
)

type HealthChecker interface {
    Healthy(string) bool
}

type HTTPHealthCheck struct {
    *http.Client
}

func NewHTTPHealthCheck(conf config.HealthCheck) *HTTPHealthCheck {
    client := &http.Client{
        Timeout: time.Duration(conf.Timeout) * time.Second,
    }

    return &HTTPHealthCheck{client}
}

func (hc *HTTPHealthCheck) Healthy(url string) (bool, error) {
    resp, err := hc.Get(url)
    if err != nil {
        return false, err
    }

    if resp.StatusCode != http.StatusOK {
        return false, nil
    }

    return true, nil
}
