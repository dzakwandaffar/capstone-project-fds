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

	authRoute := r.Group("/auth")
	authRoute.POST("/login", handler.NewAuth().Login)

	accountRoute := r.Group("/account")
	accountRoute.GET("/get", handler.NewAccount().GetAccount)
	accountRoute.POST("/create", handler.NewAccount().CreateAccount)
	accountRoute.PATCH("/update/:id", handler.NewAccount().UpdateAccount)
	accountRoute.DELETE("/delete/:id", handler.NewAccount().DeleteAccount)

	transactionRoute := r.Group("/transaction")
	transactionRoute.GET("/get", handler.NewTransaction().Get)
	transactionRoute.GET("/biller", handler.NewTransaction().GetTransaction)
	transactionRoute.GET("/biller/list-account", handler.NewTransaction().GetListAccount)
	transactionRoute.GET("/biller/:billerid/:accountid", handler.NewTransaction().GetCheckBiller)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
