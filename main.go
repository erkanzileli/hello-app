package main

import (
	"bytes"
	"io"
	"log"

	"github.com/klauspost/compress/gzip"

	"github.com/valyala/fasthttp"
)

const serverAddr = ":8081"

func main() {
	server := &fasthttp.Server{}

	server.Handler = func(ctx *fasthttp.RequestCtx) {
		body := ctx.Request.Body()
		ctx.Response.SetBody([]byte(`{"msg": "Hello!"}`))
		ctx.Response.SetStatusCode(200)
		if string(ctx.Request.Header.Peek("Content-Type")) == "gzip" {
			gzipReader, err := gzip.NewReader(bytes.NewReader(body))
			if err != nil {
				return
			}
			parsedBody, err := io.ReadAll(gzipReader)
			if err != nil {
				return
			}
			body = parsedBody
		}
		log.Printf("%s %s %s\n", ctx.Method(), ctx.Request.URI().RequestURI(), body)

	}

	log.Println("Running on", serverAddr)
	log.Fatalln(fasthttp.ListenAndServe(serverAddr, server.Handler))
}
