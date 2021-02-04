package tronics

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

var e = echo.New()

func Start() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	e.GET("/", initProduct)
	e.GET("/product", getProduct)
	e.POST("/product", createProduct)
	e.GET("/product/:id", getProductByID)
	e.PUT("/product/:id", updateProduct)
	e.DELETE("/product/:id", destroyProduct)
	e.Logger.Print(fmt.Sprintf("Listen on port %s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}
