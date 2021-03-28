package app

import (
	"bookstore_users-api/controllers/ping"
	"movies-user/controllers/users"
)

func mapUrls(){
	router.GET("/ping",ping.Ping)
	router.POST("/users",users.CreateUser)
}