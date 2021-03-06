package app

import (
	"github.com/anfelo/bookstore_oauth-api/src/clients/mongodb"
	"github.com/anfelo/bookstore_oauth-api/src/http"
	"github.com/anfelo/bookstore_oauth-api/src/repository/db"
	"github.com/anfelo/bookstore_oauth-api/src/repository/rest"
	"github.com/anfelo/bookstore_oauth-api/src/services"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication starts the web server
func StartApplication() {
	client, ctx, cancel := mongodb.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	atService := services.NewService(db.NewRepository(), rest.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8081")
}
