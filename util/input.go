package util

import "flag"

var Input = struct {
	Threads uint64
	Requests uint64
	Delay uint64

	Method  string
	Url     string
	Headers string
	Body    string

	Ca string
	Cert string
	Key string
} {}

func init() {
	flag.Uint64Var(&Input.Threads, "thN", 1, "number of threads")
	flag.Uint64Var(&Input.Requests, "rN", 1, "number of requests for every thread")
	flag.Uint64Var(&Input.Delay, "t", 1000, "sleep time after every request, unit is millisecond")

	flag.StringVar(&Input.Method, "X", "GET", "method for your request")
	flag.StringVar(&Input.Url, "url", "", "url for your request")
	flag.StringVar(&Input.Headers, "H", "", "headers for your request, , json format required")
	flag.StringVar(&Input.Body, "d", "", "body for your request")

	flag.StringVar(&Input.Ca, "ca", "", "ca cert for https request")
	flag.StringVar(&Input.Cert, "cert", "", "client certificate for https request")
	flag.StringVar(&Input.Key, "key", "", "client private certificate key for https request")

	flag.Parse()
}
