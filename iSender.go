package main
 
import (
	"sync/atomic"
	"fmt"
	"time"
	"bytes"
	"os"
	"io/ioutil"
	"encoding/json"
	"iSender/util"
)
 
var succeded uint64
var sendedRequests uint64
 
func main() {
	fmt.Println("Begin-- threads:", util.Input.Threads, ", requests for every thread:", util.Input.Requests)
	for i:=uint64(0); i<util.Input.Threads; i++ {
		go request()
	}
	for atomic.LoadUint64(&sendedRequests) < util.Input.Threads * util.Input.Requests {
		time.Sleep(time.Second)
	}
	fmt.Println("End-- sended requests:", atomic.LoadUint64(&sendedRequests), ", succed:", atomic.LoadUint64(&succeded))
}
 
func request() {
	for i:=uint64(0); i < util.Input.Requests; i++ {
		var client *util.Client
		var err error
		if util.Input.Ca == "" {
			client = util.NewClient()
		} else if util.Input.Cert == "" {
			client, err = util.NewClientWithCaPath(util.Input.Ca)
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
		response, err := client.SendRequest(request)
		if err == nil {
			fmt.Println("response code:", response.StatusCode)
			resBody, _ := ioutil.ReadAll(response.Body)
			fmt.Println("response body:",string(resBody))
			atomic.AddUint64(&succeded, 1)
		} else {
			fmt.Println("err:", err)
		}
		atomic.AddUint64(&sendedRequests, 1)
		time.Sleep(time.Duration(util.Input.Delay) * time.Millisecond)
	}
}
