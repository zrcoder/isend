# iSender
iSender is a very light tool written in Go, which can imitate to send requests with threads(infact, goroutines)<br>
It will be very helpful when you do some benchmark test for your server

## Install
you can build the source code to get the iSender binary fit for your platform. For example, type:
```
go get github.com/zrcoder/iSender
```
and then you will find the binary named "iSender" under $GOPATH/bin/<br>
now we can move it to /usr/local/bin for usage:
```
mv $GOPATH/bin/iSender /usr/local/bin
```

## Examples
```
iSender -url http://localhost:8080/test
```
```
iSender -thN 10 -url http://127.0.0.1:8080/test
```
```
iSender -rN 2 -X POST -url http://localhost:8080/test
```
```
iSender -thN 100 -rN 10 -t 1500 -X POST -H '{"Content-Type":"application/json"}' -d '{"someKey":"someValue"}' -url https://localhost:8080/test -ca ./ca.crt -cert ./client.crt -key ./client.key
```

## Usage
You can type this line for help:
```
./iSender --help
```
and you will see something like below:
```
Usage of iSender:
  -thN uint
     	number of threads (default 1)
  -rN uint
    	number of requests for each thread (default 1)
  -t uint
    	sleep time after each request, unit is millisecond (default 100)

  -H string
    	headers for your request, json format required
  -X string
    	method for your request (default "GET")
  -d string
    	body for your request
  -url string
    	url for your request

  -v
      print detail information
      
  -ca string
    	ca cert for https request
  -cert string
    	client certificate for https request
  -key string
    	client private certificate key for https request
```
