package http

import (
	"net/http"
	"strings"

	"github.com/anfelo/bookstore_oauth-api/src/domain/accesstoken"
	"github.com/gin-gonic/gin"
)

// AccessTokenHandler access token http handler interface
type AccessTokenHandler interface {
	GetByID(c *gin.Context)
}

type accessTokenHandler struct {
	service accesstoken.Service
}

// NewHandler returns a new access token http handler
func NewHandler(service accesstoken.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service,
	}
}

func (h *accessTokenHandler) GetByID(c *gin.Context) {
	accessTokenID := strings.TrimSpace(c.Param("access_token_id"))

	accessToken, err := h.service.GetByID(accessTokenID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)

	c.JSON(http.StatusNotImplemented, "implement me!")
}
