package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	hc "github.com/zrcoder/httpclient"
)

var client *hc.Client

func main() {
	client = initClient()

	fmt.Printf("Begin-- threads: %d, requests for each thread: %d\n",
		Input.Threads, Input.Requests)

	cachedchan := make(chan uint64, Input.Threads)
	for i := uint64(0); i < Input.Threads; i++ {
		go request4Thread(i, cachedchan)
	}
	// block the main thread util all the goroutines finished
	for t := uint64(0); t < Input.Threads; t += <-cachedchan {
	}

	fmt.Println("--End")
}

func request4Thread(thread uint64, c chan uint64) {
	for i := uint64(0); i < Input.Requests; i++ {
		begin := time.Now()
		response, err := client.Go()
		if err == nil {
			bodyData, _ := io.ReadAll(response.Body)
			if Input.ShowDetail {
				fmt.Printf(`[OK]
[thread %d, request %d]
<-response CODE: %d
<-response HEADER: %v
<-response BODY: %s
<-time cost: %v
`,
					thread, i, response.StatusCode, response.Header, bodyData, time.Since(begin))
			} else {
				fmt.Println("[OK] code:", response.StatusCode)
			}
		} else {
			if Input.ShowDetail {
				fmt.Printf("[Failed]\t[thread %d, request %d] error: %s\n",
					thread, i, err.Error())
			} else {
				fmt.Println("[Failed]", err)
			}
		}

		if i < Input.Requests-1 {
			time.Sleep(time.Duration(Input.Delay) * time.Millisecond)
		}
	}

	c <- 1
}

func initClient() *hc.Client {
	client := hc.New().InsecureSkipVerify(true)
	if Input.Ca != "" {
		client.AddCAFile(Input.Ca)
	}
	if Input.Cert != "" && Input.Key != "" {
		client.AddCertFile(Input.Cert, Input.Key)
	}
	if Input.Headers != "" {
		var headers map[string]string
		err := json.Unmarshal([]byte(Input.Headers), &headers)
		if err != nil {
			fmt.Println("unmarshal headers failed:", err)
			os.Exit(1)
		}
		for k, v := range headers {
			client.Header(k, v)
		}
	}
	return client.ReNew(Input.Method, Input.Url).Body(Input.Body)
}
