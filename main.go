package main

import (
	"log"
	"net/http"

	"gg"
)

func main() {
	r := gg.New()

	r.GET("/", func(c *gg.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *gg.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gg.Context) {
		c.JSON(http.StatusOK, gg.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	log.Printf("Server is running on http://localhost:9999")
	r.Run(":9999")
}
