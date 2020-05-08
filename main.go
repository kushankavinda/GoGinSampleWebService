package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kushankavinda/GoGinSampleWebService/controller"
	_ "github.com/kushankavinda/GoGinSampleWebService/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	c := controller.NewController()

	v1 := r.Group("/api/v1")
	{
		accounts := v1.Group("/accounts")
		{
			accounts.GET(":id", c.ShowAccount)
			accounts.GET("", c.ListAccounts)
			/*		accounts.POST("", c.AddAccount)
					accounts.DELETE(":id", c.DeleteAccount)
					accounts.PATCH(":id", c.UpdateAccount)
					accounts.POST(":id/images", c.UploadAccountImage) */
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
