package request

import "net/http"

type Request struct {
	url    string
	client http.Client
}

func New(url string) Request {
	return Request{url: url, client: http.Client{}}
}
