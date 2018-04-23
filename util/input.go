package util

import (
	"flag"
	"fmt"
	"os"
)

var Input = struct {
	Threads  uint64
	Requests uint64
	Delay    uint64

	Method  string
	Url     string
	Headers string
	Body    string

	Ca   string
	Cert string
	Key  string
}{}

const (
	helpInfo = `Usage of iSender:
  -thN uint
     	number of threads (default 1)
  -rN uint
    	number of requests for each thread (default 1)
  -t uint
    	sleep time after every request, unit is millisecond (default 1000)

  -H string
    	headers for your request, json format required
  -X string
    	method for your request (default "GET")
  -d string
    	body for your request
  -url string
    	url for your request

  -ca string
    	ca cert for https request
  -cert string
    	client certificate for https request
  -key string
    	client private certificate key for https request`
)

func init() {
	help := false
	h := false
	flag.BoolVar(&help, "help", false, "help info")
	flag.BoolVar(&h, "h", false, "help info")

	flag.Uint64Var(&Input.Threads, "thN", 1, "number of threads")
	flag.Uint64Var(&Input.Requests, "rN", 1, "number of requests for each thread")
	flag.Uint64Var(&Input.Delay, "t", 1000, "sleep time after every request, unit is millisecond")

	flag.StringVar(&Input.Method, "X", "GET", "method for your request")
	flag.StringVar(&Input.Url, "url", "", "url for your request")
	flag.StringVar(&Input.Headers, "H", "", "headers for your request, json format required")
	flag.StringVar(&Input.Body, "d", "", "body for your request")

	flag.StringVar(&Input.Ca, "ca", "", "ca cert for https request")
	flag.StringVar(&Input.Cert, "cert", "", "client certificate for https request")
	flag.StringVar(&Input.Key, "key", "", "client private certificate key for https request")

	flag.Parse()

	if help || h {
		fmt.Println(helpInfo)
		os.Exit(0)
	}
}
