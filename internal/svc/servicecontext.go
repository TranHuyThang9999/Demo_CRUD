package svc

import (
	"p8/internal/adapter/database"
	"p8/internal/config"
	"p8/internal/middleware"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config     config.Config
	Collection database.UserModel
	Example    rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		Collection: database.NewUserModel(conn),
		Example:    middleware.NewExampleMiddleware().Handle,
	}
}
