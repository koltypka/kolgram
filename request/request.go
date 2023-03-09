package request

import (
	"io"
	"net/http"
	"net/url"

	"github.com/koltypka/kolgram/myErr"
)

type Request struct {
	host   string
	client http.Client
}

func New(host string) Request {
	return Request{host: host, client: http.Client{}}
}

func (r *Request) run(method string, query url.Values) (data []byte, err error) {
	defer func() { err = myErr.Handle("Request error", err) }()

	myUrl := url.URL{Scheme: "https", Host: r.host, Path: method}

	req, err := http.NewRequest(http.MethodGet, myUrl.String(), nil)

	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()

	result, err := r.client.Do(req)

	if err != nil {
		return nil, err
	}

	defer func() { _ = result.Body.Close() }()

	body, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
