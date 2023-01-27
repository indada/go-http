package main

import (
	"encoding/json"
	"errors"
	"fmt"
	http2 "goxx/http"
	"log"
	"net/http"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()       // 解析参数，默认是不会解析的
	fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	/*for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}*/
	data, _ := json.Marshal(r.URL)
	fmt.Println("r.url", string(data))
}
func bb(w http.ResponseWriter, r *http.Request) {

}

type signUpReq struct {
	Aa int    `json:"aa"`
	Bb string `json:"bb"`
}
type commonResponse struct {
	Code int
	Msg  string
}

func main() {
	//m := make(map[string]string, 2)
	server := http2.NewSdkHttpServer("test-server", http2.MetricFilterBuilder, http2.MetricFilterBuilders)

	/*server.Route("/aa", sayhelloName) // 设置访问的路由
	server.Route("/bb", bb)           // 设置访问的路由*/
	server.Route(http.MethodGet, "/cc", func(ctx *http2.Context) {
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
	}) // 设置访问的路由
	server.Route(http.MethodGet, "/aa", http2.SignUp) // 设置访问的路由
	err := server.Start(":9090")                      // 设置监听的端口

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	errors.New("ess")
}
