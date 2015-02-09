package signaling

import (
	"fmt"
	"net/http"
	"net/url"
)

func Listen(urlString string) {
	parsed, parseErr := url.Parse(urlString)
	if parseErr != nil {
		panic(fmt.Sprintf("%v", parseErr))
	}
	if parsed.Scheme == "wss" {
		panic("Listen URL looks like TLS. Plese use ListenTLS function.")
	}
	http.Handle(parsed.Path, NewWebSocketListener())
	fmt.Printf("Signaling server started on \"%s\"\n", urlString)
	if err := http.ListenAndServe(parsed.Host, nil); err != nil {
		panic(fmt.Sprintf("%v", err))
	}
}

func ListenTLS(urlString, cert, key string) {
	parsed, parseErr := url.Parse(urlString)
	if parseErr != nil {
		panic(fmt.Sprintf("%v", parseErr))
	}
	if parsed.Scheme == "ws" {
		panic("Listen URL looks like not TLS. Plese use Listen function.")
	} else if cert == "" || key == "" {
		panic("Cert file or key file is empty. cannot listen TLS serve.")
	}

	http.Handle(parsed.Path, NewWebSocketListener())
	fmt.Printf("Signaling server started on \"%s\"\n", urlString)
	if err := http.ListenAndServeTLS(parsed.Host, cert, key, nil); err != nil {
		panic(fmt.Sprintf("%v", err))
	}
}
