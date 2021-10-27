package main

import (
	"net/http"

	"gee"
)

/**
* @project: go-gee
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/24/21 1:45 PM
**/

func main() {
	r := gee.New()

	r.Use(gee.Logger) // global middleware

	// curl localhost:9999/assets/backgroudimage.jpeg
	r.Static("/assets", "./static")

	// curl localhost:9999/
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v1 := r.Group("/v1")
	{
		// curl "localhost:9999/v1/hello?name=rc"
		v1.GET("/hello", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})

		//  curl -X POST -d "username=rc" -d "password=rc" localhost:9999/v1/login
		v1.POST("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	v2 := r.Group("/v2")
	v2.Use(gee.OnlyForV2)
	{
		// curl "localhost:9999/v2/hello/rc"
		v2.GET("/hello/:name", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})

		// curl localhost:9999/v2/assets/css/a.css
		v2.GET("/assets/*filepath", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
		})
	}

	r.Run(":9999")
}
