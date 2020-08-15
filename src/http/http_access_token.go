package http

import (
	"net/http"

	"github.com/anfelo/bookstore_oauth-api/src/services"
	"github.com/anfelo/bookstore_utils/errors"

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
	service services.Service
}

// NewHandler returns a new access token http handler
func NewHandler(service services.Service) AccessTokenHandler {
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
	var atReq accesstoken.AccessTokenResquest
	if err := c.ShouldBindJSON(&atReq); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	at, err := h.service.Create(atReq)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, at)
}

func (h *accessTokenHandler) UpdateExpirationTime(c *gin.Context) {

}
