package main

import (
	"sync/atomic"
	"fmt"
	"time"
	"sender/rest"
	"bytes"
	"os"
	"io/ioutil"
	"sender/input"
)

var index int64 = 0
var success int64 = 0
var actualTotalRequest int64 = 0

func main() {
	for i:=int64(0); i<input.Input.Threads; i++ {
		go request()
	}
	for atomic.LoadInt64(&actualTotalRequest) < input.Input.Requests {
		time.Sleep(time.Second)
	}
	fmt.Println("finished, actual total request:", atomic.LoadInt64(&actualTotalRequest), "success:", atomic.LoadInt64(&success))
}

func request() {
	for i:=atomic.AddInt64(&index, 1); i <= input.Input.Requests; i=atomic.AddInt64(&index, 1) {
		var client *rest.Client
		var err error
		if input.Input.Ca == "" {
			client = rest.NewClient()
		} else if input.Input.Cert == "" {
			client, err = rest.NewClientWithCaPath(input.Input.Ca)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			client, err = rest.NewClientWithCertFiles(input.Input.Ca, input.Input.Cert, input.Input.Key)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		body := bytes.NewReader([]byte(input.Input.Body))
		headers := map[string]string{
			"Content-Type": "application/json",
		}
		request, _ := rest.NewRequest(input.Input.Method, input.Input.Url, body, headers)
		response, err := client.SendRequest(request)
		if err == nil {
			fmt.Println("response code:", response.StatusCode)
			resBody, _ := ioutil.ReadAll(response.Body)
			fmt.Println("response body:",string(resBody))
			atomic.AddInt64(&success, 1)
		} else {
			fmt.Println("err:", err)
		}
		atomic.AddInt64(&actualTotalRequest, 1)
	}
}
