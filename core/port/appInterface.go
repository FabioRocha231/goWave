package port

type App interface {
	GET(path string, handler Handler, mws ...Handler)
	POST(path string, handler Handler, mws ...Handler)
	DELETE(path string, handler Handler, mws ...Handler)
	PATCH(path string, handler Handler, mws ...Handler)
	PUT(path string, handler Handler, mws ...Handler)
	Use(mw ...Handler)
	Handle(handler Handler)
	ListAndServe(address string) error
}
