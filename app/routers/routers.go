package routers

import (
	_ "assignment/docs"
	"log"

	"assignment/app/handler"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() {
	OrderHandler := handler.NewOrderHandler()
	r := gin.Default()

	r.POST("/Order", OrderHandler.AddOrder)
	r.PUT("/Order", OrderHandler.AddOrder)
	r.GET("/Order", OrderHandler.GetOrderList)
	r.GET("/Order/:order_id", OrderHandler.GetOrderDetail)
	r.DELETE("/Order/:order_id", OrderHandler.DeleteOrder)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	log.Println("=========== Server started ===========")
	r.Run(":1337")
}
