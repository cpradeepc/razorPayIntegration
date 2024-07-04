package main

import (
	"app/api"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatalln("error invoked to load env file:", err)
	}

}
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("index.html")
	api.InitRoute(router)
}
