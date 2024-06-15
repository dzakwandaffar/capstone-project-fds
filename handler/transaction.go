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
	GetListAccount(*gin.Context)
	GetCheckBiller(*gin.Context)
	Get(*gin.Context)
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

func (a *transactionImplement) Get(g *gin.Context) {
	queryParam := g.Request.URL.Query()

	start_date := queryParam.Get("start_date")

	end_date := queryParam.Get("end_date")

	transactions := []model.Transaction{}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	conds := ""
	if start_date != "" && end_date != "" {
		conds = "transaction_date BETWEEN '" + start_date + "' AND '" + end_date
	}

	result := orm.Find(&transactions, conds)

	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Get transaction successfully",
		"data":    transactions,
	})
}

type Biller struct {
	Data []struct {
		BillerID string `json:"BillerID"`
		Name     string `json:"Name"`
	} `json:"Data"`
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

type BillerListAccount struct {
	ListAccount []struct {
		BillerID  string  `json:"BillerID"`
		Amount    float64 `json:"Amount"`
		Name      string  `json:"Name"`
		AccountID string  `json:"AccountID"`
	} `json:"ListAccount"`
}

func (b *transactionImplement) GetListAccount(g *gin.Context) {

	var err error
	var client = &http.Client{}
	var data = BillerListAccount{}

	request, err := http.NewRequest("GET", "http://localhost:5000/v1/api/biller/list-account", nil)
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

type CheckBiller struct {
	Biller struct {
		BillerID  string  `json:"BillerID"`
		Amount    float64 `json:"Amount"`
		Name      string  `json:"Name"`
		AccountID string  `json:"AccountID"`
		Paid      bool
	} `json:"Biller"`
}

func (b *transactionImplement) GetCheckBiller(g *gin.Context) {

	var err error
	var client = &http.Client{}
	var data = CheckBiller{}

	billerid := g.Param("billerid")
	accountid := g.Param("accountid")

	request, err := http.NewRequest("GET", "http://localhost:5000/v1/api/biller/"+billerid+"/"+accountid, nil)
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

//================ blm selesai mutasi

// type CheckMutasi struct {
// 	Mutasi struct {
// 		BillerID  string  `json:"BillerID"`
// 		Amount    float64 `json:"Amount"`
// 		Name      string  `json:"Name"`
// 		AccountID string  `json:"AccountID"`
// 		Paid      bool
// 	} `json:"Mutasi"`
// }

// func (b *transactionImplement) GetCheckMutasi(g *gin.Context) {

// 	queryParam := g.Request.URL.Query()

// 	name := queryParam.Get("name")

// 	accounts := []model.Account{}

// 	orm := utils.NewDatabase().Orm
// 	db, _ := orm.DB()

// 	defer db.Close()

// 	result := orm.Find(&accounts, "name = ?", name)

// 	if result.Error != nil {
// 		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 			"error": result.Error,
// 		})
// 		return
// 	}

// 	g.JSON(http.StatusOK, gin.H{
// 		"message": "Get account successfully",
// 		"data":    accounts,
// 	})
// }

//================ blm selesai mutasi
