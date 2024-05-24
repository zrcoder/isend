package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags:                  flags,
		Action:                 action,
		HideHelp:               true,
		UseShortOptionHandling: true,
		Usage:                  "send requests for benchmark test",
		UsageText:              "isend -v https://test.com",
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

func action(ctx *cli.Context) error {
	if ctx.NArg() == 0 {
		cli.ShowAppHelp(ctx)
		return cli.Exit("[ERROR] url must be applied", 1)
	}

	input.Url = ctx.Args().Get(0)
	input.Method = strings.ToUpper(input.Method)

	client := newClient()

	fmt.Printf("Begin-- virtual users: %d, requests for each user: %d\n",
		input.Routines, input.Requests)

	begin := time.Now()

	wg := &sync.WaitGroup{}
	wg.Add(int(input.Routines))
	for i := uint64(0); i < input.Routines; i++ {
		go func(i uint64) {
			requestForUser(client, i)
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Println("total time:", time.Since(begin), "--End")

	return nil
}

func requestForUser(client *resty.Client, user uint64) {
	request := client.R().SetBody(input.Body).EnableTrace()
	buf := &strings.Builder{}
	for i := uint64(0); i < input.Requests; i++ {
		resp, err := request.Execute(input.Method, input.Url)
		if err != nil {
			buf.WriteString(
				fmt.Sprintf("[ERROR]\t[User %d, Request %d] error: %s\n",
					user, i, err.Error()))
			fmt.Println(buf.String())
			continue
		}

		ti := resp.Request.TraceInfo()
		buf.WriteString(fmt.Sprintf("[User %d, Request %d] %s %v\n", user, i, resp.Status(), ti.TotalTime))
		if input.HelpCount > 0 {
			writeTraceInfo(buf, &ti)
		}
		if input.HelpCount == 2 {
			buf.WriteString(fmt.Sprintf("Respons Header: %s\nResponse Body : %s\n", resp.Header(), resp.Body()))
		}
		buf.WriteString("\n")
		fmt.Println(buf.String())
	}
}

func writeTraceInfo(buf *strings.Builder, ti *resty.TraceInfo) {
	buf.WriteString(fmt.Sprintf("DNSLookup     : %v\n", ti.DNSLookup))
	buf.WriteString(fmt.Sprintf("ConnTime      : %v\n", ti.ConnTime))
	buf.WriteString(fmt.Sprintf("TCPConnTime   : %v\n", ti.TCPConnTime))
	buf.WriteString(fmt.Sprintf("TLSHandshake  : %v\n", ti.TLSHandshake))
	buf.WriteString(fmt.Sprintf("ServerTime    : %v\n", ti.ServerTime))
	buf.WriteString(fmt.Sprintf("ResponseTime  : %v\n", ti.ResponseTime))
	buf.WriteString(fmt.Sprintf("TotalTime     : %v\n", ti.TotalTime))
	buf.WriteString(fmt.Sprintf("IsConnReused  : %v\n", ti.IsConnReused))
	buf.WriteString(fmt.Sprintf("IsConnWasIdle : %v\n", ti.IsConnWasIdle))
	buf.WriteString(fmt.Sprintf("ConnIdleTime  : %v\n", ti.ConnIdleTime))
	buf.WriteString(fmt.Sprintf("RequestAttempt: %d\n", ti.RequestAttempt))
	if ti.RemoteAddr != nil {
		buf.WriteString(fmt.Sprintf("RemoteAddr    : %s\n", ti.RemoteAddr.String()))
	}
}
