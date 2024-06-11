package main

import (
	"api_gateaway/handler"
	"api_gateaway/proto"
	"context"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/client/grpc"
	micro "go-micro.dev/v4"
	"go-micro.dev/v4/client"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"*"},
	}))

	addrServiceTransactionOpt := client.WithAddress(":8084")
	clientSrvTransaction := grpc.NewClient()

	srvTransaction := micro.NewService(
		micro.Client(clientSrvTransaction),
	)

	srvTransaction.Init(
		micro.Name("service-transaction"),
		micro.Version("latest"),
	)

	accountRoute := r.Group("/account")
	accountRoute.GET("/get", handler.NewAccount().GetAccount)
	accountRoute.POST("/create", handler.NewAccount().CreateAccount)
	accountRoute.PATCH("/update/:id", handler.NewAccount().UpdateAccount)
	accountRoute.POST("/balance/", handler.NewAccount().BalanceAccount)
	accountRoute.DELETE("/delete/:id", handler.NewAccount().DeleteAccount)

	transactionRoute := r.Group("/transaction")
	transactionRoute.GET("/get", func(g *gin.Context) {
		clientResponse, err := proto.NewServiceTransactionService("service-transaction", srvTransaction.Client()).
			Login(context.Background(), &proto.LoginRequest{
				Username: "admin",
				Password: "password123",
			}, addrServiceTransactionOpt)

		if err != nil {
			g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		g.JSON(http.StatusOK, gin.H{
			"data": clientResponse,
		})
	})
	transactionRoute.POST("/transfer-bank", handler.NewTransaction().CreateTransaction)

	authRoute := r.Group("/auth")
	authRoute.POST("/login", handler.NewAuthRepository().Login)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}