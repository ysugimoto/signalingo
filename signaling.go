package signaling

import (
	"fmt"
	"github.com/ysugimoto/signalingo/connection"
	"github.com/ysugimoto/signalingo/env"
	"github.com/ysugimoto/signalingo/hooks"
	"github.com/ysugimoto/signalingo/logger"
	"github.com/ysugimoto/signalingo/operation"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

var manager *Manager

func Listen(env env.Env) {
	initLog(env)
	manager = NewManager(env)
	http.Handle(env.Server.Endpoint, createWebSocketHandler(env))
	url := fmt.Sprintf("%s:%d", env.Server.Host, env.Server.Port)
	logger.Info(fmt.Sprintf("Signaling server started on \"%s\"\n", url))
	if err := http.ListenAndServe(url, nil); err != nil {
		panic(fmt.Sprintf("%v", err))
	}
}

func ListenTLS(env env.Env) {
	initLog(env)
	manager = NewManager(env)
	http.Handle(env.Server.Endpoint, createWebSocketHandler(env))
	url := fmt.Sprintf("%s:%d", env.Server.Host, env.Server.Port)
	logger.Info(fmt.Sprintf("Signaling server started on \"%s\"\n", url))
	if err := http.ListenAndServeTLS(url, env.Server.Cert, env.Server.Key, nil); err != nil {
		panic(fmt.Sprintf("%v", err))
	}
}

func initLog(env env.Env) {
	logger.SetLevel(env.Log.Level)
	if env.Log.Type == "file" {
		logger.SetLogFile(env.Log.Filepath)
	}
}

func createWebSocketHandler(env env.Env) websocket.Handler {
	return websocket.Handler(func(ws *websocket.Conn) {
		client := connection.NewConnection(ws)

		// Handshake message
		if msg, err := operation.NewHandshakeMessage(client.UUID, manager.GetAllUsers()); err == nil {

			// Does need hook?
			if env.Hook.Url == "" {
				if err := client.Send(msg); err != nil {
					log.Println("Handshake send failed")
				} else {
					log.Printf("UUID: %s handshake", client.UUID)
				}
			} else {
				hook := hooks.NewWebHook(env.Hook.Url, client.Extra)
				hook.Run()
				resp := <-hook.Resp
				if resp == 0 {
					if err := client.Send(msg); err != nil {
						log.Println("Handshake send failed")
					} else {
						log.Printf("UUID: %s handshake", client.UUID)
					}
				} else {
					client.Close()
					return
				}
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
				log.Printf("[WebSocket] message: %s\n", msg)
				manager.HandleMessage(msg)
			}
		}

	})
}
