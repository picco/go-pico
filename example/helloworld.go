package example

import (
	"net/http"
)

// HelloWorld type
type HelloWorld struct {
	pico.Core
}

// Start func
func (app *HelloWorld) Start() {
	app.UseConfig().UseEnv().Set("APP_SERVER_ADDR", "3000")
	app.UseServer(app.Config.Get("APP_SERVER_ADDR"))

	app.Router.Get("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello world!"))
	})

	app.Server.Start()
}
