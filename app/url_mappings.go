package app

import (
	"github.com/nishant01/mybookstore_users-api/controllers/ping"
	"github.com/nishant01/mybookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	// router.GET("/users/all", users.ListUser)
	// router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
	router.PUT("/users", users.UpdateUser)
	router.DELETE("/users", users.UpdateUser)
}
