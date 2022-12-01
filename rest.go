package fluent

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type FluentRest struct {
	url     string
	headers map[string]string
	body    io.ReadCloser
	method  string
}

func GET(url string) *FluentRest {
	return &FluentRest{url: url, method: "GET"}
}

func POST(url string) *FluentRest {
	return &FluentRest{url: url, method: "POST"}
}

func (f *FluentRest) WithHeader(key string, value string) *FluentRest {
	if f.headers == nil {
		f.headers = make(map[string]string)
	}

	f.headers[key] = value

	return f
}

func (f *FluentRest) WithBody(body interface{}) *FluentRest {
	data, _ := json.Marshal(body)
	f.body = io.NopCloser(bytes.NewReader(data))

	return f
}

func (f *FluentRest) Call() *http.Response {
	req, err := http.NewRequest(f.method, f.url, f.body)

	if err != nil {
		log.Print("Error message:", err.Error())
	}

	if f.headers != nil {
		for key, value := range f.headers {
			req.Header.Set(key, value)
		}
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalln(err.Error())
	}

	return res
}
