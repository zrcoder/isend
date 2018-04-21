package input

import "flag"

var Input = Parameters{}
type Parameters struct {
	Ca string
	Cert string
	Key string
	Method string
	Header string
	Url string
	Body string
	Threads int64
	Requests int64
}

func init() {
	flag.StringVar(&Input.Ca, "ca", "", "ca cert path")
	flag.StringVar(&Input.Cert, "cert", "", "client certificate path")
	flag.StringVar(&Input.Key, "key", "", "client private certificate key path")
	flag.StringVar(&Input.Url, "url", "", "url for your request")
	flag.StringVar(&Input.Method, "X", "GET", "method for your request")
	flag.StringVar(&Input.Header, "H", "", "headers for your request")
	flag.StringVar(&Input.Body, "d", "", "body for your request")
	flag.Int64Var(&Input.Threads, "thN", 100, "the number of thread you want")
	flag.Int64Var(&Input.Requests, "rN", 10, "the number of requests for one thread you want")
	flag.Parse()
}
