package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransferInterface interface {
	TransferBank(*gin.Context)
}

type transferImplement struct{}

func Transfer() TransferInterface {
	return &transferImplement{}
}

type BodyPayloadTransfer struct{}

func (b *transferImplement) TransferBank(g *gin.Context) {

	bodyPayloadTransfer := BodyPayloadTransfer{}
	err := g.BindJSON(&bodyPayloadTransfer)

	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Hello guys this API rest for later",
	})
}