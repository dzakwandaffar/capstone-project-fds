package main

import (
	"api_gateway/handler"
	//"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	authRoute := r.Group("/auth")
	authRoute.POST("/login", handler.NewAuth().Login)

	accountRoute := r.Group("/account")
	accountRoute.POST("/balance", handler.NewAccount().CreateAccount)

	transactionRoute := r.Group("/transaction")
	transactionRoute.POST("/transfer-bank", handler.NewTransaction().TransferBank)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// import (
// 	"api_gateway/usecase"
// 	"fmt"
// )

// func main() {
// 	login := usecase.NewLogin()
// 	auth := login.Autentikasi("admin", "admin123")
// 	fmt.Println(auth)
// }
