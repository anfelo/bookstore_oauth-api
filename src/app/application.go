package app

import (
	"github.com/anfelo/bookstore_oauth-api/src/domain/accesstoken"
	"github.com/anfelo/bookstore_oauth-api/src/http"
	"github.com/anfelo/bookstore_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication starts the web server
func StartApplication() {
	atService := accesstoken.NewService(db.NewRepository())
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)

	router.Run(":8080")
}
