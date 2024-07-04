package api

import (
	"app/handler"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine) {
	r.POST("/order", handler.CreateOrder)
	//r.POST("/pay-process", handler.PaymentProcess)
	r.GET("/pay/:id", handler.CreatePayment)
	r.GET("/payment-success", handler.PaymentSuccess)
	r.GET("/payment-failed", handler.PaymentFailed)
	host := fmt.Sprintf("%s:%s", os.Getenv("hostname"), os.Getenv("port"))
	err := r.Run(host)
	if err != nil {
		log.Fatalln("server error: ", err)
	}
}
