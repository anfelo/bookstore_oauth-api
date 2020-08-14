package http

import (
	"net/http"

	"github.com/anfelo/bookstore_oauth-api/src/utils/errors"

	"github.com/anfelo/bookstore_oauth-api/src/domain/accesstoken"
	"github.com/gin-gonic/gin"
)

// AccessTokenHandler access token http handler interface
type AccessTokenHandler interface {
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	UpdateExpirationTime(c *gin.Context)
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
	accessTokenID := c.Param("access_token_id")

	accessToken, err := h.service.GetByID(accessTokenID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var at accesstoken.AccessToken
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	if err := h.service.Create(at); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, at)
}

func (h *accessTokenHandler) UpdateExpirationTime(c *gin.Context) {

}
