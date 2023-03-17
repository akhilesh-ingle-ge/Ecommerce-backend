package routes

import (
	"net/http"
	"os"

	"github.com/ecom/pkg/controller"
	"github.com/ecom/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func RunAPI() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome to Our Mini Ecommerce")
	})

	apiRoutes := router.Group("/api")

	userRoutes := apiRoutes.Group("/user")

	{
		userRoutes.POST("/register", controller.AddUser)
		userRoutes.POST("/signin", controller.SignInUser)

		// with otp
		// userRoutes.POST("/signin", controller.OtpSignInUser)
		// userRoutes.POST("/verify", controller.OtpVerify)
	}

	userProtectedRoutes := apiRoutes.Group("/users", middleware.AuthorizeJWT())
	{
		userProtectedRoutes.GET("/", controller.GetAllUsers)
		userProtectedRoutes.GET("/:user", controller.GetUser)
		userProtectedRoutes.GET("/:user/products", controller.GetProductOrdered)
		userProtectedRoutes.PUT("/:user", controller.UpdateUser)
		userProtectedRoutes.DELETE("/:user", controller.DeleteUser)
	}

	productRoutes := apiRoutes.Group("/products", middleware.AuthorizeJWT())
	{
		productRoutes.GET("/", controller.GetAllProducts)
		productRoutes.GET("/:product", controller.GetProduct)
		productRoutes.POST("/", controller.AddProduct)
		productRoutes.PUT("/:product", controller.UpdateProduct)
		productRoutes.DELETE("/:product", controller.DeleteProduct)
	}

	orderRoutes := apiRoutes.Group("/order", middleware.AuthorizeJWT())
	{
		orderRoutes.POST("/product/:product/quantity/:quantity", controller.OrderProduct)
	}

	// fileRoutes := r.Group("/file")
	// {
	// 	fileRoutes.POST("/single", handler.SingleFile)
	// 	fileRoutes.POST("/multi", handler.MultipleFile)
	// }

	router.Run(":" + os.Getenv("PORT"))
}
