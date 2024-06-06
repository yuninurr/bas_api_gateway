package main

import (
	// "apigateway/handler"
	"api_gateway/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	authRoute := r.Group("/auth")
  authRoute.POST("/login", handler.NewAutentifikasi().AutentifikasiAccount)

	accountRoute := r.Group("/account")
	accountRoute.GET("/get", handler.NewAccount().GetAccount)
	accountRoute.POST("/create", handler.NewAccount().CreateAccount)
	accountRoute.PATCH("/update/:id", handler.NewAccount().UpdateAccount)
	accountRoute.DELETE("/delete/:id", handler.NewAccount().DeleteAccount)
	accountRoute.POST("/balance", handler.NewAccount().BalanceAccount)

	transactionRoute := r.Group("/transaction")
	transactionRoute.POST("/transfer-bank", handler.Transfer().TransferBank)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// package main

// import (
// 	"api_gateway/handler"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
//   r := gin.Default()

//   // account root 
//   accountRoute := r.Group("/account")
//   accountRoute.GET("/get", handler.NewAccount().GetAccount)
//   accountRoute.POST("/create", handler.NewAccount().CreateAccount)
//   accountRoute.DELETE("/delete/:id", handler.NewAccount().DeleteAccount)
//   accountRoute.PATCH("/update/:id", handler.NewAccount().UpdateAccount)
//   accountRoute.GET("/balance", handler.NewAccount().BalanceAccount)
  
 
//   r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

// //   r.GET("/ping", func(c *gin.Context) {
// //     c.JSON(http.StatusOK, gin.H{
// //       "message": "pong",
// //     })
// //   })
//   r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
// }

// // import (
// // 	auth "api_gateway/usecase"
// // 	"fmt"
// // )

// // func main() {
	
// // 	login := auth.NewLogin()

// // 	username := "admin"
// // 	password := "admin123"

// // 	isAuthenticated := login.Authentication(username, password)

// // 	if isAuthenticated {
// // 		fmt.Println("Login successful")
// // 	} else {
// // 		fmt.Println("Invalid username or password")
// // 	}
// // }
