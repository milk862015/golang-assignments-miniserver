package module

import (
	"fmt"
	"io"
	"net/http"
)

// RunHttpServer  启动
func RunHttpServer(srv *http.Server) error {
	http.HandleFunc("/", HttpServerHandler)
	fmt.Println("[mini]run http server")
	return srv.ListenAndServe()
}

// HttpServerHandler http handler
func HttpServerHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Welcome!!!")
}
