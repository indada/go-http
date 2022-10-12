package httpv2

import (
	"fmt"
	"net/http"
)

type Server interface {
	Routable
	Start(address string) error
}

type sdkHttpServer struct {
	Name    string
	handler Handler
	root    Filter
}

//Route 路由
/*func (s *sdkHttpServer) Route(pat string, hand http.HandlerFunc) {
	http.HandleFunc(pat, hand)
}*/
func (s *sdkHttpServer) Route(method string, pat string, hand HandleFunc) {
	/*http.HandleFunc(pat, func(writer http.ResponseWriter, request *http.Request) {
		ctx := NewContext(writer, request)
		hand(ctx)
	})
	key := s.handler.keys(method, pat)
	fmt.Println(key)
	s.handler.Handlers[key] = hand*/
	s.handler.Route(method, pat, hand)
}

func (s *sdkHttpServer) Start(address string) error {
	//http.Handle("/", s.handler)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		c := NewContext(writer, request)
		s.root(c)
	})
	return http.ListenAndServe(address, nil)
}
func NewSdkHttpServer(name string, builders ...FilterBuilder) Server {
	handler := NewHandlerBasedOnMap()
	var root Filter = handler.ServeHTTP
	for i := len(builders) - 1; i >= 0; i-- {
		b := builders[i]
		root = b(root)
	}
	return &sdkHttpServer{
		Name:    name,
		handler: handler,
		root:    root,
	}
}

type signUpReq struct {
	Aa int    `json:"aa"`
	Bb string `json:"bb"`
}
type commonResponse struct {
	Code int
	Msg  string
}

func SignUp(ctx *Context) {
	req := &signUpReq{}
	err := ctx.ReadJson(req)
	if err != nil {
		_ = ctx.OkJson(&commonResponse{
			Code: 2,
			Msg:  fmt.Sprintf("INVALID REQUEST :%v", err),
		})
		return
	}
	_ = ctx.OkJson(&commonResponse{
		Code: 1,
		Msg:  fmt.Sprintf("ok"),
	})
}
