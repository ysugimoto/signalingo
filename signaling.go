package signaling

import (
	"fmt"
	"github.com/ysugimoto/signalingo/logger"
	"net/http"
)

func Listen(env Env) {
	initLog(env)
	http.Handle(env.Server.Endpoint, NewWebSocketConnectionHandler(env))
	url := fmt.Sprintf("%s:%d", env.Server.Host, env.Server.Port)
	logger.Info(fmt.Sprintf("Signaling server started on \"%s\"\n", url))
	if err := http.ListenAndServe(url, nil); err != nil {
		panic(fmt.Sprintf("%v", err))
	}
}

func ListenTLS(env Env) {
	initLog(env)
	http.Handle(env.Server.Endpoint, NewWebSocketConnectionHandler(env))
	url := fmt.Sprintf("%s:%d", env.Server.Host, env.Server.Port)
	logger.Info(fmt.Sprintf("Signaling server started on \"%s\"\n", url))
	if err := http.ListenAndServeTLS(url, env.Server.Cert, env.Server.Key, nil); err != nil {
		panic(fmt.Sprintf("%v", err))
	}
}

func initLog(env Env) {
	logger.SetLevel(env.Log.Level)
	if env.Log.Type == "file" {
		logger.SetLogFile(env.Log.Filepath)
	}
}
