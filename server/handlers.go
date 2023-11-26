package server

import (
	"net/http"
	"omni/models"

	"github.com/gin-gonic/gin"
)

func (env *Env) listUsersHandler(c *gin.Context) {
	users, err := env.users.ListUsers()

	if len(users) == 0 || err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.IndentedJSON(http.StatusOK, users)
	}
}

func (env *Env) listUserByIDHandler(c *gin.Context) {
	id := c.Param("id")

	user, err := env.users.ListUserById(id)

	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.IndentedJSON(http.StatusOK, user)
	}
}

func (env *Env) createUserHandler(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	} else {
		u, err := env.users.CreateUser(user)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
		} else {
			c.IndentedJSON(http.StatusOK, u)
		}

	}
}

func (env *Env) loginUserHandler(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	} else {
		at, err := env.auth.LoginUser(user)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
		} else {
			c.IndentedJSON(http.StatusOK, at)
		}

	}
}
