package http

import (
	"fmt"
	"go_http/core/port"
	"net/http"
)

type App struct {
	mux         *http.ServeMux
	middlewares []port.Middleware
}

func NewApp() port.App {
	return &App{
		mux:         http.NewServeMux(),
		middlewares: []port.Middleware{},
	}
}

func (app *App) GET(path string, handler port.Handler, mws ...port.Handler) {
	app.mux.HandleFunc(fmt.Sprintf("GET %s", path), app.applyMiddlewares(handler))
}

func (app *App) POST(path string, handler port.Handler, mws ...port.Handler) {
	app.mux.HandleFunc(fmt.Sprintf("POST %s", path), app.applyMiddlewares(handler))
}

func (app *App) PUT(path string, handler port.Handler, mws ...port.Handler) {
	app.mux.HandleFunc(fmt.Sprintf("PUT %s", path), app.applyMiddlewares(handler))
}

func (app *App) DELETE(path string, handler port.Handler, mws ...port.Handler) {
	app.mux.HandleFunc(fmt.Sprintf("PUT %s", path), app.applyMiddlewares(handler))
}

func (app *App) PATCH(path string, handler port.Handler, mws ...port.Handler) {
	app.mux.HandleFunc(fmt.Sprintf("GET %s", path), app.applyMiddlewares(handler))
}

func (app *App) Use(mws ...port.Middleware) {
	app.middlewares = append(app.middlewares, mws...)
}

func (app *App) applyMiddlewares(finalHandler port.Handler) port.Handler {
	// for i := len(app.middlewares) - 1; i >= 0; i-- {
	// 	finalHandler = app.middlewares[i](finalHandler)
	// }

	return finalHandler
}

func (app *App) ListAndServe(port string) error {
	fmt.Println("Running in port: ", port)

	return http.ListenAndServe(port, app.mux)
}
