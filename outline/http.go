package outline

import (
	"crypto/tls"
	"net/http"
	"time"
)

type ConnConfig struct {
	InsecureSkipVerify bool
	Timeout            time.Duration
}

func NewHttpClient(cfg ConnConfig) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: cfg.InsecureSkipVerify},
		},
		Timeout: cfg.Timeout,
	}
}
