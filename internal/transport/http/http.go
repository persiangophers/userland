package http

type HttpHandler interface {
	Start(address string) error
	Shutdown() error
}
