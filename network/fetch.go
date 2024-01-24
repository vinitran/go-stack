package network

import "net/http"

type FetchFn func(url string, init *http.Request) (*http.Response, error)

var DefaultFetchOpts = http.Request{
	Header: http.Header{
		"X-Hiro-Product": []string{"stacksjs"},
	},
}
