package main

import (
	"github.com/zrcoder/iSender/util"	
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var (
	client  = initClient()
	headers = initHeaders()
)

func main() {
	fmt.Printf("Begin-- threads: %d, requests for each thread: %d\n",
		util.Input.Threads, util.Input.Requests)

	cachedchan := make(chan uint64, util.Input.Threads)
	for i := uint64(0); i < util.Input.Threads; i++ {
		go request4Thread(i, cachedchan)
	}
	// block the main thread util all the goroutines finished
	for t := uint64(0); t < util.Input.Threads; t += <-cachedchan {
	}

	fmt.Println("--End")
}

func request4Thread(thread uint64, c chan uint64) {
	for i := uint64(0); i < util.Input.Requests; i++ {
		request := makeRequest()
		begin := time.Now()
		response, err := client.SendRequest(request)
		end := time.Now()
		if err == nil {
			if util.Input.ShowDetail {
				resBody, _ := ioutil.ReadAll(response.Body)
				fmt.Printf(`[thread %d, request %d]
<-response CODE: %d
<-response HEADER: %v
<-response BODY: %s
<-time cost: %v
`,
					thread, i, response.StatusCode, response.Header, resBody, end.Sub(begin))
			} else {
				fmt.Println("response CODE:", response.StatusCode)
			}
			response.Body.Close()
		} else {
			if util.Input.ShowDetail {
				fmt.Printf("[thread %d, request %d] error: %s\n",
					thread, i, err.Error())
			} else {
				fmt.Println("error:", err)
			}
		}
		if i < util.Input.Requests-1 {
			time.Sleep(time.Duration(util.Input.Delay) * time.Millisecond)
		}
	}
	c <- 1
}

func initClient() (client *util.Client) {
	if util.Input.Ca == "" {
		client = util.NewClient()
		return
	}
	var err error
	if util.Input.Cert == "" {
		client, err = util.NewClientWithCaFile(util.Input.Ca)
		if err != nil {
			fmt.Println("create request client with ca faild:", err)
			os.Exit(1)
		}
		return
	}
	client, err = util.NewClientWithCertFiles(util.Input.Ca, util.Input.Cert, util.Input.Key)
	if err != nil {
		fmt.Println("create request client with certificates faild:", err)
		os.Exit(1)
	}
	return client
}

func initHeaders() (headers map[string]string) {
	if util.Input.Headers != "" {
		err := json.Unmarshal([]byte(util.Input.Headers), &headers)
		if err != nil {
			fmt.Println("unmarshal headers failed:", err)
			os.Exit(1)
		}
	}
	return
}

// request can't be reused, because after sended, it's body will be clened by http package~
// so we must make a new request for every send.
func makeRequest() (request *http.Request) {
	body := bytes.NewReader([]byte(util.Input.Body))
	request, err := util.NewRequest(util.Input.Method, util.Input.Url, body, headers)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}
