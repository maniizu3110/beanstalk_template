package handler

import (
	"net/http"
	"server/api/service"
	"server/api/sqlc"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func AssignUserHandler(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			store := c.Get("store").(sqlc.Store)
			s := service.NewUserService(store)

			c.Set("Service", s)
			return handler(c)
		}
	})
	// g.GET("/:id", GetUserHandler)
	g.GET("/", ListUserHandler)
}

// func GetUserHandler(c echo.Context) error {
// 	service := c.Get("Service").(service.UserService)
// 	type getUserParams struct {
// 		Username string
// 	}
// 	params := &getUserParams{}
// 	c.Bind(params)
// 	User, err := service.GetUser(params.Username)
// 	if err != nil {
// 		logrus.Info(err.Error())
// 		return err
// 	}
// 	return c.JSON(http.StatusOK, User)
// }

func ListUserHandler(c echo.Context) error {
	service := c.Get("Service").(service.UserService)
	//TODO:offsetとlimitをoptionにする
	params := sqlc.ListUserParams{}
	c.Bind(params)
	if params.Limit == 0 {
		params.Limit = 10
	}
	users, err := service.ListUser(params)
	if err != nil {
		logrus.Info(err.Error())
		return err
	}
	return c.JSON(http.StatusOK, users)

}
