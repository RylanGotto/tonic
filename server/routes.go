package server

import "github.com/gin-gonic/gin"

func (env *Env) userRoutes(r *gin.Engine) {
	r.GET("/users", env.listUsersHandler)
	r.GET("/users/:id", env.listUserByIDHandler)
	r.POST("/users", env.createUserHandler)
}

func (env *Env) authRoutes(r *gin.Engine) {
	r.POST("/auth/login", env.loginUserHandler)
	// r.POST("/auth/logout", env.logoutUser)
}
