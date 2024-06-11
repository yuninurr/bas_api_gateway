package handler

import (
	"api_gateaway/models"
	"api_gateaway/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionInterface interface {
	TransferBank(*gin.Context)
	CreateTransaction(*gin.Context)
}

type transactionImplement struct{}

func NewTransaction() transactionInterface {
	return &transactionImplement{}
}

// type bodyPayloadTransaction struct{}

func (a *transactionImplement) TransferBank(g *gin.Context) {
	bodyPayload := models.Transaction{}

	err := g.BindJSON(&bodyPayload)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
	}

	// timeNow := time.Now()
	// bodyPayload.TransactionDate = &timeNow

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	result := orm.Create(&bodyPayload)
	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Get account successfully",
		"data":    bodyPayload,
	})
}

func (a *transactionImplement) CreateTransaction(g *gin.Context) {
	bodyPayload := models.Transaction{}

	err := g.BindJSON(&bodyPayload)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
	}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	result := orm.Create(&bodyPayload)
	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": "400",
			"error": result.Error,
		})
	} else {
		g.JSON(http.StatusOK, gin.H{
			"code": "200",
			"message": "Berhasil Membuat Data Transaksi",
			"data": bodyPayload,
		})
	}
}