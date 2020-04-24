package app

import (
	"github.com/26keisuke/go_demo/controllers/ping"
	"github.com/26keisuke/go_demo/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
}