# iSender
iSender is a very light tool writed with Golang, which can imitate to send requests with threads<br>
It will be very helpful when you do some benchmark test for your server

## Usage
Download [iSender](./iSender), place it in some proper directoey(eg. /tmp)<br>
Firstly, we must add executable permissions for iSender
```
chmod +x iSender
```
You can type this line for help:
```
./iSender --help
```
and you will see something like below:
```
  -H string
    	headers for your request
  -X string
    	method for your request (default "GET")
  -ca string
    	ca cert for https request
  -cert string
    	client certificate for https request
  -d string
    	body for your request
  -key string
    	client private certificate key for https request
  -rN uint
    	number of requests for every thread (default 1)
  -t uint
    	sleep time after every request, unit is millisecond (default 1000)
  -thN uint
    	number of threads (default 1)
  -url string
    	url for your request
```
## Examples
```
./iSender -url https://github.com
```
```
./iSender -thN 10 -url https://github.com
```
```
./iSender -rN 2 -X POST -url https://github.com
```
```
threads=100
requests=10
delay=1500
heads='{"Content-Type":"application/json"}'
body='{"someKey":"someValue"}'
method=POST
url=https://localhost:8080/test
ca=./ca.crt
cert=./client.crt
key=./client.key
./iSender -thN $threads -rN $requests -t $delay -H $headers -d $body -X $method -url $url -ca $ca -cert $cert -key $key
```