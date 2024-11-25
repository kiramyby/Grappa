package grappa

import (
	"log"
	"time"
)

func Logger() HandleFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
