package hello

import (
	"fmt"
	"io"
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
)

// This example shows the minimal code needed to get a restful.WebService working.
//
// GET http://localhost:8080/hello

func TryHello() {
	ws := new(restful.WebService)
	fmt.Println("server started")
	ws.Route(ws.GET("/hello").To(hello))
	ws.Route(ws.GET("/secret").Filter(basicAuthenticate).To(secret))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "world")
}

func basicAuthenticate(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	// usr/pwd = admin/admin
	u, p, ok := req.Request.BasicAuth()
	if !ok || u != "admin" || p != "admin" {
		resp.AddHeader("WWW-Authenticate", "Basic realm=Protected Area")
		resp.WriteErrorString(401, "401: Not Authorized")
		return
	}
	chain.ProcessFilter(req, resp)
}

func secret(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "Successful auth")
}
