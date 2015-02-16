package static

import (
	"github.com/ysugimoto/husky"
	"path/filepath"
)

func Serve(host string, port int) {
	app := husky.NewApp()
	app.Config.Set("host", host)
	app.Config.Set("port", port)
	app.Config.Set("path", "/")

	app.Get("/", func(d *husky.Dispatcher) {
		url := d.Input.GetRequest().URL.Path
		if resource, err := Asset("public" + url); err != nil {
			d.Output.SetHeader("Content-Type", "text/plain")
			d.Output.SetStatus(404)
			d.Output.Send(err.Error())
		} else {
			mime := findMimeType(filepath.Ext(url))
			d.Output.SetHeader("Content-Type", mime)
			d.Output.SetStatus(200)
			d.Output.Send(resource)
		}
	})
	app.Get("/playground", func(d *husky.Dispatcher) {
		if html, err := Asset("public/playground.html"); err != nil {
			d.Output.SetHeader("Content-Type", "text/plain")
			d.Output.SetStatus(404)
			d.Output.Send(err.Error())
		} else {
			d.Output.SetHeader("Content-Type", "text/html")
			d.Output.SetStatus(200)
			d.Output.Send(html)
		}
	})

	app.Serve()
}

func findMimeType(ext string) (mime string) {
	switch ext {
	case ".css":
		mime = "text/css"
	case ".js":
		mime = "text/javascript"
	case ".json":
		mime = "application/json"
	case ".gif":
		mime = "image/gif"
	case ".jpg", ".jpeg":
		mime = "image/jpeg"
	case ".png":
		mime = "image/png"
	}
	return
}
