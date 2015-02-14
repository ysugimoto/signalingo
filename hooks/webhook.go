package hooks

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type WebHook struct {
	Url    string
	Fields string
	Resp   chan int
}

func NewWebHook(url string, extra string) *WebHook {
	return &WebHook{
		Url:    url,
		Fields: extra,
		Resp:   make(chan int, 1),
	}
}

func (h *WebHook) Run() {
	request, err := h.createRequest()
	if err != nil {
		h.Resp <- 1
		return
	}

	response, fail := h.sendRequest(request)
	if fail != nil {
		h.Resp <- 1
		return
	}

	defer response.Body.Close()

	log.Printf("WebHook status code: %d\n", response.StatusCode)

	if response.StatusCode == 200 {
		h.Resp <- 0
	} else {
		h.Resp <- 1
	}
}

func (h *WebHook) sendRequest(request *http.Request) (*http.Response, error) {
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

	return client.Do(request)
}

func (h *WebHook) createRequest() (*http.Request, error) {
	request, err := http.NewRequest("POST", h.Url, strings.NewReader(h.Fields))
	if err != nil {
		log.Printf("WebHook::createRequest request create error: %v", err)
		return nil, err
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return request, nil
}
