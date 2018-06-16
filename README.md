# iSender
iSender is a very light tool written in Golang, which can imitate to send requests with threads<br>
It will be very helpful when you do some benchmark test for your server

## Download
[iSender for Linux](./bin/linux/iSender)<br>

or you can build the source code to get the iSender binary fit for your platform. For example, type:
```
go get github.com/DingHub/iSender
```
and then you will find the binary named "iSender" under $GOPATH/bin/

## Examples
```
./iSender -url http://localhost:8080/test
```
```
./iSender -thN 10 -url http://127.0.0.1:8080/test
```
```
./iSender -rN 2 -X POST -url http://localhost:8080/test
```
```
./iSender -thN 100 -rN 10 -t 1500 -X POST -H '{"Content-Type":"application/json"}' -d '{"someKey":"someValue"}' -url https://localhost:8080/test -ca ./ca.crt -cert ./client.crt -key ./client.key
```

## Usage
Firstly, we must add executable permission for iSender
```
chmod +x iSender
```
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
    	sleep time after each request, unit is millisecond (default 1000)

  -H string
    	headers for your request, json format required
  -X string
    	method for your request (default "GET")
  -d string
    	body for your request
  -url string
    	url for your request

  -v bool
      print detail information
      
  -ca string
    	ca cert for https request
  -cert string
    	client certificate for https request
  -key string
    	client private certificate key for https request
```
