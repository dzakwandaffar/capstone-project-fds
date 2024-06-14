package main

import (
	"api_gateway/handler"

	//"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	//"go-micro.dev/v4/cmd/protoc-gen-micro/plugin/micro"
	//"google.golang.org/grpc"
	//"google.golang.org/grpc/profiling/service"
	//"go-micro.dev/v4/util/ctx"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"*"},
	}))

	authRoute := r.Group("/v1")
	authRoute.POST("/login", handler.NewAuth().Login)

	accountRoute := r.Group("/account")
	accountRoute.GET("/get", handler.NewAccount().GetAccount)
	accountRoute.POST("/create", handler.NewAccount().CreateAccount)
	accountRoute.PATCH("/update/:id", handler.NewAccount().UpdateAccount)
	accountRoute.DELETE("/delete/:id", handler.NewAccount().DeleteAccount)
	accountRoute.POST("/balance", handler.NewAccount().BalanceAccount)

	transactionRoute := r.Group("/transaction")
	transactionRoute.POST("/transfer-bank", handler.NewTransaction().TransferBank)

	// transactionRoute := r.Group("/")
	// transactionRoute.POST("/transfer-bank", handler.NewTransaction().TransferBank)
	transactionRoute.GET("/biller", handler.NewTransaction().GetTransaction)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
