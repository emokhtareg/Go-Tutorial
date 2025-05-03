package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/fasthttp/router"
	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"
)

var port = ":8080"

func main() {

	fmt.Println("Server starting on port 8080")
	if false {
		basic()
		basicWithMux()
		basicChi()
		basicFasthttp()
		basicFiber()
		basicEcho()
	}
	basicGinGonic()

}

func basic() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Basic demo")
	})
	log.Fatal(http.ListenAndServe(port, nil))

}

func basicWithMux() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "basicWithMux demo")
	})
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "about : basicWithMux demo")
	})

	log.Fatal(http.ListenAndServe(port, mux))

}

func basicChi() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("basicChi demo"))
	})
	http.ListenAndServe(port, r)
}

func basicFasthttp() {
	r := router.New()
	r.GET("/", func(ctx *fasthttp.RequestCtx) {
		ctx.WriteString("Welcome!")
	})
	r.GET("/hello/{name}", func(ctx *fasthttp.RequestCtx) {
		fmt.Fprintf(ctx, "Hello, %s!\n", ctx.UserValue("name"))
	})

	log.Fatal(fasthttp.ListenAndServe(port, r.Handler))
}

func basicFiber() {
	app := fiber.New()

	app.Static("/", "./static")

	// Respond with "Hello, World!" on root path, "/"
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.SendString("Hello, basicFiber!")
	})

	app.Listen(port)
}

func basicEcho() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, basicEcho!")
	})
	e.Logger.Fatal(e.Start(port))
}

func basicGinGonic() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		 
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(port) // listen and serve on 0.0.0.0:8080
}
