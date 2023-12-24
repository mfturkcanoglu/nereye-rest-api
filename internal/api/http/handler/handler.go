package handler

import "github.com/mfturkcan/nereye-rest-api/internal/api/http/server"

type Handler interface {
	RegisterRoutes(router *server.CustomRouter)
}
