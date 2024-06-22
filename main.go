package main

import (
	"log"
	"os"

	"github.com/codeabuu/ecom/controllers"
	"github.com/codeabuu/ecom/database"
	"github.com/codeabuu/ecom/middleware"
	"github.com/codeabuu/ecom/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.User(gin.Logger())

	routes.UserRoutes(router)
	router.User(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}
