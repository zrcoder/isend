package util

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"crypto/x509"
	"crypto/tls"
	"io"
	"github.com/DingHub/iSender/util/cache"
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

func NewClientWithCaFile(ca string) (*Client, error) {
	if caContent, found := certsCache.Search(ca); found {
		if caBytes, ok := caContent.([]byte); ok {
			return NewClinetWithCaContent(caBytes)
		}
		return nil, fmt.Errorf("content in chache is not []byte")
	}
	caContent, err := ioutil.ReadFile(ca)
	if err != nil {
		return nil, err
	}
	certsCache.Add(ca, caContent)
	return NewClinetWithCaContent(caContent)
}

func NewClinetWithCaContent(caContent []byte) (*Client, error) {
	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(caContent) {
		return nil, fmt.Errorf("append cert from pem failed");
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			RootCAs:            pool,
		},
	}
	client := Client{&http.Client{Transport: tr}}
	return &client, nil
}

func NewClientWithCertFiles(ca, cert, key string) (*Client, error) {
	var caContent, certContent, keyContent cache.T
	var err error
	var found bool

	if caContent, found = certsCache.Search(ca); !found {
		caContent, err = ioutil.ReadFile(ca)
		if err != nil {
			return nil, fmt.Errorf("Read ca cert failed, %s", err.Error())
		}
		certsCache.Add(ca, caContent)
	}
	if certContent, found = certsCache.Search(cert); !found {
		certContent, err = ioutil.ReadFile(cert)
		if err != nil {
			return nil, err
		}
		certsCache.Add(cert, certContent)
	}
	if keyContent, found = certsCache.Search(key); !found {
		keyContent, err = ioutil.ReadFile(key)
		if err != nil {
			return nil, err
		}
		certsCache.Add(key, keyContent)
	}

	caBytes, ok := caContent.([]byte)
	if !ok {
		return nil, fmt.Errorf("content in chache is not []byte")
	}
	certBytes, ok := certContent.([]byte)
	if !ok {
		return nil, fmt.Errorf("content in chache is not []byte")
	}
	keyBytes, ok := keyContent.([]byte)
	if !ok {
		return nil, fmt.Errorf("content in chache is not []byte")
	}
	return NewClientWithCertsContent(caBytes, certBytes, keyBytes)
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
			InsecureSkipVerify: true,
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
