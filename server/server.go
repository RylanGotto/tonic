package server

import (
	"database/sql"
	"omni/models"

	"github.com/gin-gonic/gin"
)

type Env struct {
	users interface {
		ListUsers() ([]models.User, error)
		ListUserById(id string) (models.User, error)
		CreateUser(usr models.User) (models.User, error)
	}
}

func Serv(db *sql.DB) {
	r := gin.Default()

	env := &Env{
		users: models.UserModel{DB: db},
	}

	r.GET("/users", env.listUsersHandler)
	r.GET("/users/:id", env.listUserByIDHandler)
	r.POST("/users", env.createUserHandler)

	r.Run() // listen and serve on
}
