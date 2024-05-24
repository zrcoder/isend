package main

import (
	"github.com/urfave/cli/v2"
)

var input struct {
	Method    string
	Url       string
	Headers   cli.StringSlice
	Body      string
	Ca        string
	Cert      string
	Key       string
	Routines  uint64
	Requests  uint64
	HelpCount int
}

var flags = []cli.Flag{
	&cli.Uint64Flag{
		Name:        "vus",
		Usage:       "virtual users `uint`",
		Value:       1,
		Destination: &input.Routines,
		Action: func(ctx *cli.Context, u uint64) error {
			if u == 0 {
				return cli.Exit("virtual users should > 0", 1)
			}
			return nil
		},
		Category: "1. COMMON",
	},
	&cli.Uint64Flag{
		Name:        "num",
		Usage:       "number of requests for each user `uint`",
		Value:       1,
		Destination: &input.Requests,
		Action: func(ctx *cli.Context, u uint64) error {
			if u == 0 {
				return cli.Exit("requests should > 0", 1)
			}
			return nil
		},
		Category: "1. COMMON",
	},
	&cli.StringFlag{
		Name:        "X",
		Aliases:     []string{"method"},
		Usage:       "method `string`",
		Value:       "GET",
		Destination: &input.Method,
		Category:    "2. REQUEST",
	},
	&cli.StringSliceFlag{
		Name: "H",
		// Aliases:     []string{"header"},
		Destination: &input.Headers,
		Usage:       "headers like `'key: value'`",
		Category:    "2. REQUEST",
	},
	&cli.StringFlag{
		Name:        "d",
		Aliases:     []string{"body"},
		Destination: &input.Body,
		Usage:       "body `string`",
		Category:    "2. REQUEST",
	},
	&cli.StringFlag{
		Name:        "ca",
		Destination: &input.Ca,
		Usage:       "ca cert `FILE` for https request",
		Category:    "3. CERTIFICATES",
	},
	&cli.StringFlag{
		Name:        "cert",
		Destination: &input.Cert,
		Usage:       "client certificate `FILE` for https request",
		Category:    "3. CERTIFICATES",
	},
	&cli.StringFlag{
		Name:        "key",
		Destination: &input.Key,
		Usage:       "client private certificate key `FILE` for https request",
		Category:    "3. CERTIFICATES",
	},
	&cli.BoolFlag{
		Name:  "v",
		Usage: "print detail information",
		Count: &input.HelpCount,
	},
}
