package handlers

func (server *Server) home() {
	// ini untuk munculkan halaman homepage
	server.Router.GET("/", RenderHomeWithAntrian)
}

func (server *Server) register() {
	// ini untuk munculkan halaman registrasi
	server.Router.GET("/register", RenderRegister)

	// ini untuk handle registrasi ketika submit form dengan method POST
	server.Router.POST("/register", HandleRegistrasi)
}

func (server *Server) apiStatus() {
	// ini untuk munculkan status API
	server.Router.GET("/api", ApiStatus)
}
