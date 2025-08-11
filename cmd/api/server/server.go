package server

import (
	"log"
	"net/http"

	"time"

	v1 "github.com/amorindev/headless-ecomm-cms/cmd/api/server/v1"
)

type HttpServer struct {
	server *http.Server
}

func NewHttpServer(port string) *HttpServer {
	apiV1 := v1.New()

	serv := &http.Server{
		Addr:         ":" + port,
		Handler:      apiV1,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := HttpServer{server: serv}
	return &server
}

func (serv *HttpServer) Start() {
	log.Printf("Http server running http://localhost%s\n", serv.server.Addr)
	log.Fatal(serv.server.ListenAndServe())
}
