package api

import (
	"server/api/handler"
	"server/api/util/lib"
	"server/api/util/middleware"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func (server *Server) SetRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.LoggingMiddleware)
	validator, err := lib.NewValidator()
	if err != nil {
		logrus.Fatal("バリデージョンの設定に失敗しました")
	}

	if server.config.Env != "prod" {
		e.Debug = true
	}
	e.Validator = validator
	middleware.CORS(e)
	{
		g := e.Group("/v1",
			middleware.SetStore(server.store),
			middleware.SetConfig(server.config),
			middleware.SetTokenMaker(server.tokenMaker),
		)
		handler.AssignSignHandler(g.Group("/sign"))
		handler.AssignUserHandler(g.Group("/user"))

	}
	return e
}

