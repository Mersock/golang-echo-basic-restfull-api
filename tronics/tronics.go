package tronics

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
)

var e = echo.New()

func init() {
	err := cleanenv.ReadEnv(&cfg)
	fmt.Printf("%+v", cfg)
	if err != nil {
		e.Logger.Fatal("Unable to load Configuration")
	}
}

func serverMsg(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("inside middleware")
		return next(c)
	}
}

func Start() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	e.Use(serverMsg)
	e.GET("/", initProduct, serverMsg)
	e.GET("/product", getProduct)
	e.POST("/product", createProduct)
	e.GET("/product/:id", getProductByID)
	e.PUT("/product/:id", updateProduct)
	e.DELETE("/product/:id", destroyProduct)
	e.Logger.Print(fmt.Sprintf("Listen on port %s", cfg.Port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", cfg.Port)))
}
