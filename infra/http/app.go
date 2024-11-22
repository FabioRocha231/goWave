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

// GET adds a route to the app that listens for GET requests to the given path.
// The request will be handled by the given handler.
// The handler will be wrapped by the given middleware functions, if any.
func (app *App) GET(path string, handler port.Handler, mws ...port.Handler) {
	app.mux.HandleFunc(fmt.Sprintf("GET %s", path), handler)
}

// POST adds a route to the app that listens for POST requests to the given path.
// The request will be handled by the given handler.
// The handler will be wrapped by the given middleware functions, if any.
func (app *App) POST(path string, handler port.Handler, mws ...port.Handler) {
	app.mux.HandleFunc(fmt.Sprintf("POST %s", path), handler)
}

// PUT adds a route to the app that listens for PUT requests to the given path.
// The request will be handled by the given handler.
// The handler will be wrapped by the given middleware functions, if any.
func (app *App) PUT(path string, handler port.Handler, mws ...port.Handler) {
	app.mux.HandleFunc(fmt.Sprintf("PUT %s", path), handler)
}

// DELETE adds a route to the app that listens for DELETE requests to the given path.
// The request will be handled by the given handler.
// The handler will be wrapped by the given middleware functions, if any.
func (app *App) DELETE(path string, handler port.Handler, mws ...port.Handler) {
	app.mux.HandleFunc(fmt.Sprintf("PUT %s", path), handler)
}

// PATCH adds a route to the app that listens for PATCH requests to the given path.
// The request will be handled by the given handler.
// The handler will be wrapped by the given middleware functions, if any.
func (app *App) PATCH(path string, handler port.Handler, mws ...port.Handler) {
	app.mux.HandleFunc(fmt.Sprintf("GET %s", path), handler)
}

// Use adds the given middleware functions to the app's list of middleware.
// The middleware functions will be applied to all handlers in the order they
// are given.
func (app *App) Use(mws ...port.Middleware) {
	app.middlewares = append(app.middlewares, mws...)
}

// func (app *App) applyMiddlewares(middlewares  port.Handler) port.Handler {
// 	// var finalHandler []port.Middleware
// 	// for i := len(app.middlewares) - 1; i >= 0; i-- {
// 	// 	finalHandler = app.middlewares[i](finalHandler)
// 	// }
// 	return middlewares
// }

// ListAndServe starts the app listening for incoming requests on the given port.
// Any middleware functions added to the app will be applied to all handlers in
// the order they were added.
func (app *App) ListAndServe(port string) error {
	fmt.Println("Running in port: ", port)

	return http.ListenAndServe(port, app.mux)
}
