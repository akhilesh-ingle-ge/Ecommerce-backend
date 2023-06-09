package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ecom/pkg/config"
	"github.com/ecom/pkg/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password *string) {
	bytePassword := []byte(*password)
	hPassword, _ := bcrypt.GenerateFromPassword(bytePassword, 10)
	*password = string(hPassword)
}

func ComparePassword(dbPass, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass)) == nil
}

func GetAllUsers(c *gin.Context) {
	var users []model.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	// intId, err := strconv.Atoi(id)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"Error": err.Error(),
	// 	})
	// 	return
	// }
	var user model.User
	config.DB.First(&user, id)
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func SignInUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	var dbUser model.User
	config.DB.First(&dbUser, "email = ?", user.Email)

	if isTrue := ComparePassword(dbUser.Password, user.Password); isTrue {
		fmt.Println("user before", dbUser.ID)
		token := GenerateToken(dbUser.ID)
		c.JSON(http.StatusOK, gin.H{
			"message": "Successfully SignIN",
			"token":   token,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{
		"Error": "Password not matched",
	})
	return
}

func AddUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	HashPassword(&user.Password)
	config.DB.Create(&user)
	user.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	var user model.User
	config.DB.First(&user, intId)
	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	config.DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	config.DB.Where("id = ?", intId).Delete(&model.User{})
	c.JSON(http.StatusOK, gin.H{
		"Message": "Success",
	})
}

func GetProductOrdered(c *gin.Context) {
	userstr := c.Param("user")
	userId, err := strconv.Atoi(userstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	var orders []model.Order
	config.DB.Where("user_id = ?", userId).Set("gorm:auto_preload", true).Find(&orders)

	c.JSON(http.StatusOK, gin.H{
		"data": orders,
	})
}
