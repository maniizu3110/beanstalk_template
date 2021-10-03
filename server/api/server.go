package api

import (
	"fmt"
	"server/api/sqlc"
	"server/api/util/lib"
	"server/api/util/token"
	"github.com/labstack/echo/v4"
)

type Server struct {
	config     lib.Config
	store      sqlc.Store
	tokenMaker token.Maker
	router     *echo.Echo
}

func NewServer(config lib.Config, store sqlc.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	server.router = server.SetRouter()
	return server, nil

}
func (server *Server) Start(address string) {
	server.router.Logger.Fatal(server.router.Start(address))
}
