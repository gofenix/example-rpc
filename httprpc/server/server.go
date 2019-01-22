package main

import (
	"example-rpc/jsonrpc/service"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.RegisterName("HelloService", new(service.HelloService))

	http.HandleFunc("/jsonrpc", func(writer http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body, Writer: writer,
		}

		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":1234", nil)
}
