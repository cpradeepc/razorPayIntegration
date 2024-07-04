package handler

import (
	"app/model"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	rpay "github.com/razorpay/razorpay-go"
)

var pageData *model.PageVariables

const (
	// RazorpayKeyID     string = "rzp_test_vOsKKSWnOE803Q"
	// RazorpayKeySecret string = "JINdUUpdybhJ707mAu37fH84"
	//apiURL           string = "https://api.razorpay.com/v1/orders"
	statusPaid       string = "paid"
	statusUnpaid     string = "unpaid"
	statusProcessing string = "processing"
)

// To create order id via razorpay.
func CreateOrder(c *gin.Context) {
	log.Printf("page data>> %+v\n ", pageData)
	err := c.ShouldBindJSON(&pageData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": nil, "error": err})
		return
	}
	pageData.UsrKey = os.Getenv("RazorpayKeyID")
	pageData.Status = statusUnpaid
	log.Printf("order data>> %+v\n ", pageData)

	//To create client
	client := rpay.NewClient(os.Getenv("RazorpayKeyID"), os.Getenv("RazorpayKeySecret"))

	//To get order id pass  mapping data
	data := map[string]interface{}{
		"amount":   pageData.Amount,
		"currency": pageData.Currency,
		"receipt":  pageData.Receipt,
	}
	body, err := client.Order.Create(data, nil)
	log.Printf("resp of order process: %+v\n", body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": nil, "error": err})
		return
	}
	ordId, ok := body["id"]
	initDate := body["created_at"]
	if ok {
		pageData.OrderId = ordId.(string)
		pageData.Created_At = initDate.(float64)
	}

	// log.Println("resp body: ", body)
	// c.JSON(http.StatusOK, gin.H{"data": *pageData, "error": nil})

	pageData.Status = statusProcessing
	//object pageData save in db server...
	//return html page url for create-payment as given order details
	payUrl := c.Request.Host + "/pay/" + pageData.OrderId
	c.JSON(http.StatusOK, gin.H{"payUrl": payUrl})
}

// create a payment according to order id as parameter via the route
func CreatePayment(c *gin.Context) {
	ord_id := c.Param("id")
	log.Println("pageData: ", pageData)
	if pageData != nil && ord_id != "" {
		if pageData.OrderId == ord_id && pageData.Status == statusPaid {
			c.JSON(http.StatusConflict, gin.H{"error": "given id payment already success"})
			return
		}
	}
	if ord_id == "" || ord_id != "" && pageData.OrderId != ord_id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "given id invalid"})
		return
	}
	c.HTML(http.StatusOK, "index.html", *pageData)
}

// If a payment may  be failed update data in db server,
func PaymentFailed(c *gin.Context) {
	ord_id := c.Query("resp_ord_id")
	pay_id := c.Query("resp_pay_id")
	code := c.Query("resp_code")
	desp := c.Query("resp_desp")
	source := c.Query("resp_source")
	step := c.Query("resp_step")
	reason := c.Query("resp_reason")

	var payFailed model.PayFailed
	payFailed.OrderId = ord_id
	payFailed.PayementId = pay_id
	payFailed.Code = code
	payFailed.Description = desp
	payFailed.Source = source
	payFailed.Step = step
	payFailed.Reason = reason

	pageData.Status = statusUnpaid
	log.Println("pageData: ", pageData)
	log.Println("Pay failed: ", payFailed, "ord_id", ord_id)

	c.JSON(http.StatusTemporaryRedirect, gin.H{"error stub": payFailed})

}

// on payment success update data in db server,
func PaymentSuccess(c *gin.Context) {
	var paySuc model.PaySuccess
	log.Println("raw >>:", c.Request.URL.RawPath, ":", c.Request.URL.RawQuery, ":", c.Request.URL.Query())
	pay_id := c.Query("pay_id")
	order_id := c.Query("ord_id")
	signature := c.Query("signature")

	// err := helper.SignatureVarify(signature, pageData.OrderId, pageData.PayId)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"data": nil, "error": err})
	// 	return
	// }

	paySuc.PaymentId = pay_id
	paySuc.OrderId = order_id
	paySuc.Signature = signature

	pageData.PayId = pay_id
	pageData.Status = statusPaid
	log.Println("paySuc data", paySuc)

	c.JSON(http.StatusOK, gin.H{"status": "success", "payment_id": pay_id, "error": nil})

}
