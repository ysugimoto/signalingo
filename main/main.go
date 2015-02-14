package main

import (
	"fmt"
	"github.com/ysugimoto/go-cliargs"
	"github.com/ysugimoto/husky"
	"github.com/ysugimoto/signalingo"
	"github.com/ysugimoto/signalingo/env"
	"io/ioutil"
	"os"
)

func main() {
	env := env.InitEnv("")
	args := cliarg.NewArguments()
	args.Alias("h", "host", env.Server.Host)
	args.Alias("p", "port", env.Server.Port)
	args.Alias("e", "endpoint", env.Server.Endpoint)
	args.Parse()

	host, _ := args.GetOptionAsString("host")
	port, _ := args.GetOptionAsInt("port")
	endpoint, _ := args.GetOptionAsString("endpoint")

	env.Server.Host = host
	env.Server.Port = port
	env.Server.Endpoint = endpoint

	if env.Server.Tls {
		go signaling.ListenTLS(env)
	} else {
		go signaling.Listen(env)
	}

	//signaling.StaticServe()

	// testing http
	app := husky.NewApp()
	app.Config.Set("host", "127.0.0.1")
	app.Config.Set("port", 54321)
	app.Config.Set("path", "/")

	app.Get("/", func(d *husky.Dispatcher) {
		pwd, _ := os.Getwd()
		d.Output.SetHeader("Content-Type", "text/html")
		html := pwd + "/public/index.html"
		if _, err := os.Stat(html); err == nil {
			buf, _ := ioutil.ReadFile(html)
			d.Output.SetStatus(200)

			d.Output.Send(buf)
		} else {
			d.Output.SetStatus(200)

			d.Output.Send(fmt.Sprintf("%v", err))
		}
	})
	app.Serve()
}
