package signaling

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Hook struct {
	Url    string
	Fields map[string]string
	Resp   chan int
}

func NewHook(url string, extra map[string]string) *Hook {
	return &Hook{
		Url:    url,
		Fields: extra,
		Resp:   make(chan int, 1),
	}
}

func (h *Hook) Run() {
	values := url.Values{}
	for key, value := range h.Fields {
		values.Add(key, value)
	}

	request, err := http.NewRequest("POST", h.Url, strings.NewReader(values.Encode()))
	if err != nil {
		h.Resp <- 1
		return
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	var client *http.Client
	parse, _ := url.Parse(h.Url)
	if parse.Scheme == "https" {
		transport := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = &http.Client{Transport: transport}
	} else {
		client = &http.Client{}
	}

	response, fail := client.Do(request)
	if fail != nil {
		h.Resp <- 1
		return
	}

	defer response.Body.Close()

	log.Printf("Hook status code: %d\n", response.StatusCode)

	if response.StatusCode == 200 {
		h.Resp <- 0
	} else {
		h.Resp <- 1
	}
}
