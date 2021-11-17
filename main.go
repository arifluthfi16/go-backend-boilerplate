package main

import (
	"fmt"
	"github.com/arifluthfi16/gomvcboilerplate/config"
	"github.com/arifluthfi16/gomvcboilerplate/routers"
	"github.com/arifluthfi16/gomvcboilerplate/services"
	"github.com/arifluthfi16/gomvcboilerplate/services/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	db     *gorm.DB
	Router *gin.Engine
}

func (server *Server) InitServer() {
	server.Router = gin.Default()
	var Router = routers.RouteLoader{}
	for _, routes := range Router.LoadRoutes() {
		routes.Route(server.Router)
	}
}

func (server *Server) InjectDBToService() {
	services.InjectDBIntoServices(server.db)
}

func (server *Server) Run() {
	fmt.Println("Rise and shine! 🌞🌞🌞")
	fmt.Println(config.ServerConfig.AppConfig.AppName + " is listening on port : 5050")
	server.Router.Run(":5050")
}

func main() {
	// Load Config from Env
	config.LoadConfig()

	// Uncomment this and remove Server{} to use database
	app := Server{db: db.LoadDB(config.ServerConfig.DBConfig)}
	app.InjectDBToService()

	app.InitServer()
	app.Run()
}
