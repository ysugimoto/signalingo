package signaling

import (
	"fmt"
	"net/http"
)

func Listen(host string, port int, endpoint string) {
	http.Handle(endpoint, NewWebSocketConnectionHandler())
	url := fmt.Sprintf("%s:%d", host, port)
	fmt.Printf("Signaling server started on \"%s\"\n", url)
	if err := http.ListenAndServe(url, nil); err != nil {
		panic(fmt.Sprintf("%v", err))
	}
}

func ListenTLS(host string, port int, endpoint, cert, key string) {
	if cert == "" || key == "" {
		panic("Cert file or key file is empty. cannot listen TLS serve.")
	}
	http.Handle(endpoint, NewWebSocketConnectionHandler())
	url := fmt.Sprintf("%s:%d", host, port)
	fmt.Printf("Signaling server started on \"%s\"\n", url)
	if err := http.ListenAndServeTLS(url, cert, key, nil); err != nil {
		panic(fmt.Sprintf("%v", err))
	}
}
