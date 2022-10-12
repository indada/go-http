package http

import (
	"fmt"
	"net/http"
)

type Routable interface {
	Route(method string, pat string, hand func(c *Context))
}
type Handler interface {
	//http.Handler
	ServeHTTP(c *Context)
	Routable
}

type HandlerBasedOnMap struct {
	Handlers map[string]func(ctx *Context)
}

func (h *HandlerBasedOnMap) Route(method string, pat string, hand func(ctx *Context)) {
	/*http.HandleFunc(pat, func(writer http.ResponseWriter, request *http.Request) {
		ctx := NewContext(writer, request)
		hand(ctx)
	})*/
	key := h.keys(method, pat)
	fmt.Println(key)
	h.Handlers[key] = hand
}
func (h *HandlerBasedOnMap) ServeHTTP(c *Context) {

	key := h.key(c.R)
	if handler, ok := h.Handlers[key]; ok {
		handler(c)
	} else {
		c.W.WriteHeader(http.StatusNotFound)
		c.W.Write([]byte("404 Not Found!!!"))
	}

}
func (h *HandlerBasedOnMap) key(req *http.Request) string {
	return req.Method + "#" + req.URL.Path
	//return req.Method + "#" + req.URL.Path
}
func (h *HandlerBasedOnMap) keys(method string, pattern string) string {
	return method + "#" + pattern
	//return req.Method + "#" + req.URL.Path
}

var _ Handler = &HandlerBasedOnMap{}

func NewHandlerBasedOnMap() Handler {
	return &HandlerBasedOnMap{
		Handlers: make(map[string]func(ctx *Context)),
	}
}
