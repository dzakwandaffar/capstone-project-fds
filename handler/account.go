package handler

import (
	"api_gateway/model"
	"api_gateway/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountInterface interface {
	GetAccount(*gin.Context)
	CreateAccount(*gin.Context)
	UpdateAccount(*gin.Context)
	DeleteAccount(*gin.Context)
	BalanceAccount(*gin.Context)
}

type accountImplement struct{}

func NewAccount() AccountInterface {
	return &accountImplement{}
}

func (a *accountImplement) GetAccount(g *gin.Context) {
	queryParam := g.Request.URL.Query()

	name := queryParam.Get("name")

	accounts := []model.Account{}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	result := orm.Find(&accounts, "name = ?", name)

	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Get account successfully",
		"data":    accounts,
	})
}

type BodyPayloadAccount struct {
}

func (a *accountImplement) CreateAccount(g *gin.Context) {
	bodyPayload := []model.Account{}

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
			"error": result.Error,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Create account successfully",
		"data":    bodyPayload,
	})
}

func (a *accountImplement) UpdateAccount(g *gin.Context) {
	// queryParam := g.Request.URL.Query()

	// name := queryParam.Get("name")
	bodyPayload := model.Account{}

	err := g.BindJSON(&bodyPayload)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}

	id := g.Param("id")

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	user := model.Account{}

	orm.First(&user, "account_id = ?", id)
	if user.AccountID == "" {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Data Not Found",
		})
		return
	}

	user.Name = bodyPayload.Name
	user.Username = bodyPayload.Username
	orm.Save(user)

	g.JSON(http.StatusOK, gin.H{
		"message": "Update account successfully",
		"data":    user,
	})
}

func (a *accountImplement) DeleteAccount(g *gin.Context) {
	id := g.Param("id")

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	result := orm.Where("account_id = ?", id).Delete(&model.Account{})
	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Delete Account successfully",
		"data":    id,
	})
}

type BodyPayloadBalance struct {
	Account_ID string
	Month      string
}

func (a *accountImplement) BalanceAccount(g *gin.Context) {
	bodyPayloadBal := BodyPayloadBalance{}

	err := g.BindJSON(&bodyPayloadBal)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
	}

	sumResult := struct {
		Total int
	}{}
	transaction := []model.Transaction{}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	q := orm
	result := q.Find(&transaction)
	orm.Model(&model.Transaction{}).Select("sum(amount) as total").Where("account_id = ? AND date_part( 'Month' , transaction_date) = ?", bodyPayloadBal.Account_ID, bodyPayloadBal.Month).Group("account_id").Scan(&sumResult)

	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Hello guys this API rest for later",
		"data":    transaction,
	})
}
