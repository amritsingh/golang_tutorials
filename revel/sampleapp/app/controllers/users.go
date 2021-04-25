package controllers

import (
	"net/http"
	"strconv"
	"sampleapp/app/models"

	"github.com/revel/revel"
)


type Users struct {
	*revel.Controller
}

func (c Users) GetUser() revel.Result {
	id, _ := strconv.ParseInt(c.Params.Route.Get("id"), 10, 0)
	user := models.User{}.GetByID(uint64(id))
	return c.RenderJSON(user)
}

type UserSt struct {
	Name    string                    `json:"name"`
	Email   string                    `json:"email"`
}

func (c Users) Create() revel.Result {
	var userParams UserSt
	c.Params.BindJSON(&userParams)
	if userParams.Email != "" {
		user := models.User{}.UpdateOrCreate(userParams.Name, userParams.Email)
		if user.ID > 0 {
			c.Response.Status = http.StatusCreated
		}
	} else {
		c.Response.Status = http.StatusBadRequest
	}
	return c.RenderText("")
}

func (c Users) Delete() revel.Result {
	id, _ := strconv.ParseInt(c.Params.Route.Get("id"), 10, 0)
	userDeleted := models.User{}.DeleteUser(uint64(id))
	if userDeleted {
		c.Response.Status = http.StatusAccepted
	} else {
		c.Response.Status = http.StatusBadRequest
	}
	return c.RenderText("")
}
