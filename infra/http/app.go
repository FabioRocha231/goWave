package http

import (
	"fmt"
	"go_http/core/port"
	"net/http"
)

type App struct {
	mux         *http.ServeMux
	middlewares []port.Handler
}

func NewApp() port.App {
	return &App{
		mux:         http.NewServeMux(),
		middlewares: []port.Handler{},
	}
}

func (app *App) GET(path string, handler port.Handler, mws ...port.Handler) {
	http.HandleFunc(fmt.Sprintf("GET %s", path), handler)
}

func (app *App) POST(path string, handler port.Handler, mws ...port.Handler) {
	http.HandleFunc(fmt.Sprintf("POST %s", path), handler)
}

func (app *App) PUT(path string, handler port.Handler, mws ...port.Handler) {
	http.HandleFunc(fmt.Sprintf("PUT %s", path), handler)
}

func (app *App) DELETE(path string, handler port.Handler, mws ...port.Handler) {
	http.HandleFunc(fmt.Sprintf("DELETE %s", path), handler)
}

func (app *App) PATCH(path string, handler port.Handler, mws ...port.Handler) {
	http.HandleFunc(fmt.Sprintf("PATCH %s", path), handler)
}

func (app *App) Use(mws ...port.Handler) {
	app.middlewares = append(app.middlewares, mws...)
}

func (app *App) Handle(handler port.Handler) {
}

func (app *App) ListAndServe(port string) error {
	fmt.Println("Running in port: ", port)

	return http.ListenAndServe(port, app.mux)
}
