package gee

import (
	"log"
	"time"
)

/**
* @project: go-gee
*
* @description:
*
* @author: rachaelyychen
*
* @create: 10/26/21 8:28 PM
**/

func Logger(c *Context) {
	// Start timer
	t := time.Now()
	// Process request
	c.Next()
	// Calculate resolution time
	log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
}

func OnlyForV2(c *Context) {
	// Start timer
	t := time.Now()
	// Process request
	c.Next()
	// Calculate resolution time
	log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
}
