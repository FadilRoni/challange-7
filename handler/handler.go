package handler

import "challange-7/service"

type HttpServer struct {
	app service.ServiceInterface
}

func NewHttpServer(app service.ServiceInterface) HttpServer {
	return HttpServer{app: app}
}
