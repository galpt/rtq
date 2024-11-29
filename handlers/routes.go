package handlers

func (server *Server) home() {
	// ini untuk munculkan halaman homepage
	server.Router.GET("/", RenderHome)
}

func (server *Server) login() {
	// ini untuk munculkan halaman login
	server.Router.GET("/login", RenderLogin)

	// ini untuk handle login ketika submit form dengan method POST
	server.Router.POST("/login", RenderLogin)
}

func (server *Server) register() {
	// ini untuk munculkan halaman registrasi
	server.Router.GET("/register", RenderRegister)

	// ini untuk handle registrasi ketika submit form dengan method POST
	server.Router.POST("/register", HandleRegistrasi)
}

func (server *Server) guest() {
	// ini untuk munculkan halaman guest
	server.Router.GET("/guest", RenderGuest)
}

func (server *Server) history() {
	// ini untuk munculkan halaman history
	server.Router.GET("/history", RenderHistory)
}

func (server *Server) profile() {
	// ini untuk munculkan halaman profile
	server.Router.GET("/profile", RenderProfile)
}

func (server *Server) security() {
	// ini untuk munculkan halaman security
	server.Router.GET("/security", RenderSecurity)
}

func (server *Server) settings() {
	// ini untuk munculkan halaman settings
	server.Router.GET("/settings", RenderSettings)
}

func (server *Server) apiStatus() {
	// ini untuk munculkan status API
	server.Router.GET("/api", ApiStatus)
}
