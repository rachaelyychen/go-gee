package gee

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
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

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		c.Next()
	}
}

// print stack trace for debug
func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:]) // skip first 3 caller

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}
