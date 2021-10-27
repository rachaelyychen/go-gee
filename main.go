package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

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

	// global middleware
	r.Use(gee.Logger)
	r.Use(gee.Recovery())

	// curl localhost:9999/assets/image/backgroundimage.png
	r.Static("/assets", "./static")

	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")

	// localhost:9999/
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	// curl localhost:9999/panic
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gee.Context) {
		var names []string
		c.String(http.StatusOK, names[1])
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

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}
