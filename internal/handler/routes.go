// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"p8/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: loginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/users/id/:userId",
				Handler: GetUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/users/create",
				Handler: CreateUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/id",
				Handler: GetByIdHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
