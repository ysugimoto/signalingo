package signaling

import (
	"fmt"
	"github.com/ysugimoto/signalingo/connection"
	"github.com/ysugimoto/signalingo/env"
	"github.com/ysugimoto/signalingo/hooks"
	"github.com/ysugimoto/signalingo/logger"
	"github.com/ysugimoto/signalingo/operation"
	"golang.org/x/net/websocket"
	"net/http"
)

var manager *Manager

func Listen(env env.Env) {
	url := initHandler(env)
	if err := http.ListenAndServe(url, nil); err != nil {
		panic(fmt.Sprintf("%v", err))
	}
}

func ListenTLS(env env.Env) {
	url := initHandler(env)
	if err := http.ListenAndServeTLS(url, env.Server.Cert, env.Server.Key, nil); err != nil {
		panic(fmt.Sprintf("%v", err))
	}
}

func initHandler(env env.Env) (url string) {
	initLog(env)
	manager = NewManager(env)
	http.Handle(env.Server.Endpoint, createWebSocketHandler(env, false))
	http.Handle("/manage", createWebSocketHandler(env, true))
	url = fmt.Sprintf("%s:%d", env.Server.Host, env.Server.Port)
	logger.Infof("Signaling server started on \"%s\"\n", url)

	return
}

func initLog(env env.Env) {
	logger.SetLevel(env.Log.Level)
	if env.Log.Type == "file" {
		logger.SetLogFile(env.Log.Filepath)
	}
}

func GracefulShutdown() {
	logger.Info("Signal received, graceful shutdown")
	if msg, err := operation.NewShutdownMessage(); err == nil {
		manager.Broadcast(msg)
	}
	manager.Purge()
}

func createWebSocketHandler(env env.Env, admin bool) websocket.Handler {
	return websocket.Handler(func(ws *websocket.Conn) {
		client := connection.NewConnection(ws, admin)

		// Handshake message
		if msg, err := operation.NewHandshakeMessage(client.UUID, manager.GetAllUsers()); err == nil {

			var resp int
			// Do we need hook?
			if env.Hook.Url == "" {
				resp = 0
			} else {
				hook := hooks.NewWebHook(env.Hook.Url, client.Extra)
				hook.Run()
				resp = <-hook.Resp
			}

			if resp == 0 {
				if err := client.Send(msg); err != nil {
					logger.Fatal("Handshake send failed")
				} else {
					logger.Infof("UUID: %s handshake", client.UUID)
				}
			} else {
				client.Close()
				return
			}
		}

		manager.AddConnection(client)

		defer func() {
			manager.CloseConnection(client)
		}()

		for {
			if client.Closed {
				break
			}
			if msg, err := client.Receive(); err != nil {
				break
			} else {
				logger.Infof("WebSocket receive message: %s\n", msg)
				manager.HandleMessage(msg)
			}
		}
	})
}
