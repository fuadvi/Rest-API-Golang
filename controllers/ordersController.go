package controllers

import (
	"assement_02/database"
	"assement_02/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Item struct {
	Item_code   uint   `json:"Item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
}

type OrderData struct {
	OrderedAt     string `json:"ordered_at"`
	CoustumerName string `json:"coustumer_name"`
	Items         []Item `json:"items"`
}

func CreateOrder(ctx *gin.Context) {
	var orders OrderData
	db := database.GetDB()

	if err := ctx.ShouldBindJSON(&orders); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var Items models.Items

	if len(orders.Items) < 1 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "item order tidak ada",
		})
	} else {
		for _, item := range orders.Items {
			Items = models.Items{
				Item_code:   item.Item_code,
				Description: item.Description,
				Quantity:    item.Quantity,
			}
		}
	}

	if result := db.Create(&Items); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "berhasil order",
		"data":    orders,
	})

}
