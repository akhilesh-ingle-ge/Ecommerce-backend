package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ecom/pkg/config"
	"github.com/ecom/pkg/model"
	"github.com/gin-gonic/gin"
)

func OrderProduct(c *gin.Context) {
	prodIdstr := c.Param("product")
	prodId, err := strconv.Atoi(prodIdstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
	} else {
		quantityIdstr := c.Param("quantity")
		quantityId, err := strconv.Atoi(quantityIdstr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
		} else {
			fmt.Println(c)
			userIdfloat := c.GetFloat64("userID")
			fmt.Println(userIdfloat)
			userId := int(userIdfloat)
			fmt.Println(userId)
			order := model.Order{
				ProductID: uint(prodId),
				UserID:    uint(userId),
				Quantity:  quantityId,
			}
			config.DB.Preload("User", "Product").Create(&order)

			// M-2
			// config.DB.Preload(clause.Associations).Create(&model.Order{
			// 	ProductID: uint(prodId),
			// 	UserID:    uint(userId),
			// 	Quantity:  quantityId,
			// })

			c.String(http.StatusOK, "Product Ordered Successfully")
		}
	}
}
