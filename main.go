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

	// curl localhost:9999/
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	// curl "localhost:9999/hello?name=rc"
	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	//  curl -X POST -d "username=rc" -d "password=rc" localhost:9999/login
	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	// curl "localhost:9999/hello/rc"
	r.GET("/hello/:name", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	// curl localhost:9999/assets/css/a.css
	r.GET("/assets/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
