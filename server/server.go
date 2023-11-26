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
	auth interface {
		LoginUser(u models.User) (models.Auth, error)
	}
}

func Serv(db *sql.DB) {
	r := gin.Default()

	r.Use(IsAuthorized())

	um := models.UserModel{DB: db}
	env := &Env{
		users: um,
		auth:  models.AuthModel{UM: um},
	}

	env.userRoutes(r)
	env.authRoutes(r)

	r.Run() // listen and serve on
}
