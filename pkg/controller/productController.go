package controller

import (
	"net/http"
	"strconv"

	"github.com/ecom/pkg/config"
	"github.com/ecom/pkg/model"
	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {
	var products []model.Product
	config.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

func GetProduct(c *gin.Context) {
	prodstr := c.Param("product")
	prodId, err := strconv.Atoi(prodstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	var product model.Product
	config.DB.First(&product, prodId)
	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

func AddProduct(c *gin.Context) {
	var product model.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	config.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

func UpdateProduct(c *gin.Context) {
	prodstr := c.Param("product")
	prodId, err := strconv.Atoi(prodstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	var product model.Product
	config.DB.First(&product, prodId)
	err = c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	config.DB.Save(&product)
	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

func DeleteProduct(c *gin.Context) {
	prodstr := c.Param("product")
	prodId, err := strconv.Atoi(prodstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	config.DB.Where("id = ?", prodId).Delete(&model.Product{})
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}
