package util

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"crypto/x509"
	"crypto/tls"
	"io"
	"iSender/util/cache"
)

type Client struct {
	client *http.Client
}

var certsCache = cache.NewWithCapacity(3)

func NewClient() *Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := Client{&http.Client{Transport: tr}}
	return &client
}

func NewClientWithCaPath(caPath string) (*Client, error) {
	if caContent, found := certsCache.Search(caPath); found {
		return NewClinetWithCaContent(caContent)
	}
	caContent, err := ioutil.ReadFile(caPath)
	if err != nil {
		return nil, fmt.Errorf("read ca cert failed, %s", err.Error())
	}
	certsCache.Add(caPath, caContent)
	return NewClinetWithCaContent(caContent)
}

func NewClinetWithCaContent(caContent []byte) (*Client, error)  {
	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(caContent) {
		return nil, fmt.Errorf("append cert from pem failed");
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
			RootCAs:            pool,
		},
	}
	client := Client{&http.Client{Transport: tr}}
	return &client, nil
}

func NewClientWithCertFiles(caPath, certPath, keyPath string) (*Client, error) {
	var caContent, certContent, keyContent []byte
	var err error
	var found bool

	if caContent, found = certsCache.Search(caPath); !found {
		caContent, err = ioutil.ReadFile(caPath)
		if err != nil {
			return nil, fmt.Errorf("Read ca cert failed, %s", err.Error())
		}
		certsCache.Add(caPath, caContent)
	}
	if certContent, found = certsCache.Search(certPath); !found {
		certContent, err = ioutil.ReadFile(certPath)
		if err != nil {
			return nil, fmt.Errorf("read client cert failed, %s", err.Error())
		}
		certsCache.Add(certPath, certContent)
	}
	if keyContent, found = certsCache.Search(keyPath); !found {
		keyContent, err = ioutil.ReadFile(keyPath)
		if err != nil {
			return nil, fmt.Errorf("read client key failed, %s", err.Error())
		}
		certsCache.Add(keyPath, keyContent)
	}
	return NewClientWithCertsContent(caContent, certContent, keyContent)
}

func NewClientWithCertsContent(caContent, certContent, keyContent []byte) (*Client, error) {
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caContent)
	certPair, err := tls.X509KeyPair(certContent, keyContent)
	if err != nil {
		return nil, fmt.Errorf("Load x509 key pair failed, %s", err.Error())
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates:       []tls.Certificate{certPair},
			ClientAuth:         tls.RequireAndVerifyClientCert,
			InsecureSkipVerify: false,
			RootCAs:            pool,
		},
	}
	Client := Client{&http.Client{Transport: tr}}
	return &Client, nil
}

func NewRequest(method string, url string, body io.Reader, headers map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if headers != nil {
		for headerName, value := range headers {
			req.Header.Set(headerName, value)
		}
	}
	return req, nil
}

func (c *Client) SendRequest(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}
