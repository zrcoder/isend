# isend
isend is a light tool written in Go, which can imitate to send requests with threads(infact, goroutines)

It will be very helpful when you do some benchmark test for your server

## Install

```
go install github.com/zrcoder/isend@latest
```

## Examples

```
isend  http://localhost:8080/test
```

```
isend -vus 10 http://127.0.0.1:8080/test
```

```
isend -nun 2 -X POST -url http://localhost:8080/test
```

```
isend -vus 100 -num 10 -X POST -H 'Content-Type: application/json' -d '{"someKey":"someValue"}' -ca ca.crt -cert client.crt -key client.key https://localhost:8080/test 
```

## Usage
Type `isend` for help, and you will see something like below:

```
NAME:
   isend - send requests for benchmark test

USAGE:
   isend -v https://test.com

GLOBAL OPTIONS:
   -v  print detail information (default: false)

   1. COMMON

   --num uint  number of requests for each user uint (default: 1)
   --vus uint  virtual users uint (default: 1)

   2. REQUEST

   -H 'key: value' [ -H 'key: value' ]  headers like 'key: value'
   -X string, --method string           method string (default: "GET")
   -d string, --body string             body string

   3. CERTIFICATES

   --ca FILE    ca cert FILE for https request
   --cert FILE  client certificate FILE for https request
   --key FILE   client private certificate key FILE for https request
```
