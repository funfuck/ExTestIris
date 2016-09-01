package controllers

import (
	"github.com/kataras/iris"
	"testIris/models"
)

type UserApi struct {
	*iris.Context
}

// GET /users
func (u UserApi) Get() {
	u.Write("Get from /users")
	u.JSON(iris.StatusOK, models.User{Name:"test"})
	// u.JSON(iris.StatusOK,myDb.AllUsers())
}

// GET /:param1 which its value passed to the id argument
func (u UserApi) GetBy(id string) { // id equals to u.Param("param1")
	u.Write("Get from /users/%s", id)
	// u.JSON(iris.StatusOK, myDb.GetUserById(id))
}

// PUT /users
func (u UserApi) Put() {
	name := u.FormValue("name")
	// myDb.InsertUser(...)
	println(string(name))
	println("Put from /users")
}

// POST /users/:param1
func (u UserApi) PostBy(id string) {

	name := u.PostValue("name") // you can still use the whole Context's features!
	// myDb.UpdateUser(...)
	println(string(name))
	println("Post from /users/" + id)
}

// DELETE /users/:param1
func (u UserApi) DeleteBy(id string) {
	// myDb.DeleteUser(id)
	println("Delete from /" + id)
}
