package main

import (
	"sync/atomic"
	"fmt"
	"time"
	"bytes"
	"os"
	"io/ioutil"
	"encoding/json"
	"github.com/DingHub/iSender/util"
)
 
var succeeded uint64
var sended uint64
 
func main() {
	fmt.Println("Begin-- threads:", util.Input.Threads, ", requests for each thread:", util.Input.Requests)
	for i := uint64(0); i < util.Input.Threads; i++ {
		go request4Thread(i)
	}
	for atomic.LoadUint64(&sended) < util.Input.Threads*util.Input.Requests {
		time.Sleep(time.Second)
	}
	fmt.Println("End-- sended requests:", atomic.LoadUint64(&sended), ", succeed:", atomic.LoadUint64(&succeeded))
}
 
func request4Thread(thread uint64) {
	var client *util.Client
	var err error
	if util.Input.Ca == "" {
		client = util.NewClient()
	} else if util.Input.Cert == "" {
		client, err = util.NewClientWithCaFile(util.Input.Ca)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		client, err = util.NewClientWithCertFiles(util.Input.Ca, util.Input.Cert, util.Input.Key)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	body := bytes.NewReader([]byte(util.Input.Body))
	var headers map[string]string
	if util.Input.Headers != "" {
		err = json.Unmarshal([]byte(util.Input.Headers), &headers)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	request, err := util.NewRequest(util.Input.Method, util.Input.Url, body, headers)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := uint64(0); i < util.Input.Requests; i++ {
		response, err := client.SendRequest(request)
		if err == nil {
			if util.Input.ShowDetail {
				resBody, _ := ioutil.ReadAll(response.Body)
				fmt.Printf("[thread %d, request %d] response CODE: %d; respnese BODY: %s\n", thread, i, response.StatusCode, resBody)
			} else {
				fmt.Println("response CODE:", response.StatusCode)
			}
			atomic.AddUint64(&succeeded, 1)
		} else {
			if util.Input.ShowDetail {
				fmt.Printf("[thread %d, request %d] error: %s\n", thread, i, err.Error())
			} else {
				fmt.Println("error:", err)
			}
		}
		atomic.AddUint64(&sended, 1)
		time.Sleep(time.Duration(util.Input.Delay) * time.Millisecond)
	}
}
