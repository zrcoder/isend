package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	DefaultTimeout               = 60 * time.Second
	DefaultDialTimeout           = 30 * time.Second
	DefaultKeepAliveTimeout      = 30 * time.Second
	DefaultIdleConnTimeout       = 90 * time.Second
	DefaultTLSHandshakeTimeout   = 10 * time.Second
	DefaultExpectContinueTimeout = 1 * time.Second
)

func newClient() *resty.Client {
	client := resty.New().
		SetTransport(&http.Transport{
			Proxy:                 http.ProxyFromEnvironment,
			IdleConnTimeout:       DefaultIdleConnTimeout,
			TLSHandshakeTimeout:   DefaultTLSHandshakeTimeout,
			ExpectContinueTimeout: DefaultExpectContinueTimeout,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}).
		SetTimeout(DefaultTimeout)

	if input.Ca != "" {
		client.SetRootCertificate(input.Ca)
	}
	if input.Cert != "" && input.Key != "" {
		cert, err := tls.LoadX509KeyPair(input.Cert, input.Key)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		client.SetCertificates(cert)
	}
	for _, header := range input.Headers.Value() {
		i := strings.Index(header, ":")
		if i == -1 {
			fmt.Println("invalid header format")
			os.Exit(1)
		}
		client.SetHeader(strings.TrimSpace(header[:i]), strings.TrimSpace(header[i+1:]))
	}
	return client
}
