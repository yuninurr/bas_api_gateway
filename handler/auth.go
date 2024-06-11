package handler

import (
	"api_gateaway/model"
	"api_gateaway/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authImplement struct{}

type authInterface interface {
	Login(*gin.Context)
}

func NewAuthRepository() authInterface {
	return &authImplement{}
}

type BodyPayloadAuth struct {
	Username string
	Password string
}

func (a *authImplement) Login(g *gin.Context) {
	bodyPayloadAuth := BodyPayloadAuth{}

	err := g.BindJSON(&bodyPayloadAuth)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	accounts := []model.Account{}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	q := orm

	if bodyPayloadAuth.Username != "" && bodyPayloadAuth.Password != "" {
		q = q.Where("username =? AND password =?", bodyPayloadAuth.Username, bodyPayloadAuth.Password)
	}

	result := q.Find(&accounts)

	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	if len(accounts) == 0 {
		g.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": "401",
			"message": "Gagal Login!, username atau password salah!",
		})
		return
	} else {
		g.JSON(http.StatusOK, gin.H{
			"code": "200",
			"message": "Berhasil Login",
			"username": bodyPayloadAuth.Username,
		})
	}

	// err := g.BindJSON(&bodyPayloadAuth)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(bodyPayloadAuth.Username, bodyPayloadAuth.Password)

	// untuk merubah response API yang dibuat (/auth/login)
	// if usecase.NewLogin().Authenticate(bodyPayloadAuth.Username, bodyPayloadAuth.Password) {
	// 	g.JSON(http.StatusOK, gin.H{
	// 		"responseCode":    "200",
	// 		"responseMessage": "anda berhasil login", "data": bodyPayloadAuth,
	// 	})
	// } else {
	// 	g.JSON(http.StatusBadRequest, gin.H{
	// 		"responseCode": "401",
	// 		"message":      "anda gagal login", "data": err,
	// 	})
	// }
}