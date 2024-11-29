package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Server ... struct to hold global variables
type Server struct {
	Router *gin.Engine
}

// Init ... Initializes the app
func (server *Server) Init(port string) {
	gin.SetMode(gin.ReleaseMode)
	server.Router = gin.New()

	// disini untuk routing custom sesuai kebutuhan
	server.home()
	server.login()
	server.register()
	server.guest()
	server.history()
	server.profile()
	server.security()
	server.settings()
	server.apiStatus()

	// Load HTML and Static files
	server.Router.LoadHTMLGlob("views/*.html")
	server.Router.Static("/css", "views/css")
	server.Router.Static("/fonts", "views/fonts")
	server.Router.Static("/img", "views/img")
	server.Router.Static("/js", "views/js")

	notify := fmt.Sprintf("Server started.\nAccess on:\n - http://0.0.0.0%v/\n - http://127.0.0.1%v/", port, port)
	fmt.Println(notify)
	server.Router.Run(port)
}
