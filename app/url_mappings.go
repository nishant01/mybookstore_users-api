package app

import (
	"github.com/nishant01/mybookstore_users-api/controllers/ping"
	"github.com/nishant01/mybookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.Get)
	// router.GET("/users/all", users.ListUser)
	router.POST("/users", users.Create)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)

	router.GET("/internal/users/search", users.Search)
}
