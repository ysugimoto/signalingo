package signaling

import (
	"fmt"
	"net/http"
)

func Listen(env Env) {
	http.Handle(env.Server.Endpoint, NewWebSocketConnectionHandler())
	url := fmt.Sprintf("%s:%d", env.Server.Host, env.Server.Port)
	fmt.Printf("Signaling server started on \"%s\"\n", url)
	if err := http.ListenAndServe(url, nil); err != nil {
		panic(fmt.Sprintf("%v", err))
	}
}

func ListenTLS(env Env) {
	http.Handle(env.Server.Endpoint, NewWebSocketConnectionHandler())
	url := fmt.Sprintf("%s:%d", env.Server.Host, env.Server.Port)
	fmt.Printf("Signaling server started on \"%s\"\n", url)
	if err := http.ListenAndServeTLS(url, env.Server.Cert, env.Server.Key, nil); err != nil {
		panic(fmt.Sprintf("%v", err))
	}
}
