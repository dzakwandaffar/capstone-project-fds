package handler

import (
	"api_gateway/model"
	"api_gateway/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionInterface interface {
	TransferBank(*gin.Context)
	GetTransaction(*gin.Context)
}

type transactionImplement struct{}

func NewTransaction() TransactionInterface {
	return &transactionImplement{}
}

type BodyPayloadTransaction struct{}

func (b *transactionImplement) TransferBank(g *gin.Context) {

	bodyPayloadTxn := model.Transaction{}
	err := g.BindJSON(&bodyPayloadTxn)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
	}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	result := orm.Create(&bodyPayloadTxn)
	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Transaction Succesfull",
		"data":    bodyPayloadTxn,
	})

}

type Biller struct {
	Data []struct {
		BillerID string `json:"BillerID"`
		Name     string `json:"Name"`
	} `json:"data"`
}

func (b *transactionImplement) GetTransaction(g *gin.Context) {

	var err error
	var client = &http.Client{}
	var data = Biller{}

	request, err := http.NewRequest("GET", "http://localhost:5000/v1/api/biller", nil)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response, err := client.Do(request)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer response.Body.Close()

	fmt.Println(response.Body)

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, data)

}
