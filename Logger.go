package middleware

import (
	"fmt"

	"github.com/varunmbet/goblack"

	"io"
	"log"
	"os"
	"time"
)

// Logger returns a baa middleware for log http access
func Logger() goblack.HandlerFunc {
	return func(c *goblack.Context) {
		start := time.Now()

		c.Next()

		c.App().Logger().Printf("%s %s %s %v %v\n", c.RemoteAddr(), c.Req.Method, c.URL(true), c.Resp.Status(), time.Since(start))
	}
}

// LoggerWithWriter returns a baa middleware with writer for log http access
func LoggerWithWriter(f io.Writer) goblack.HandlerFunc {
	return func(c *goblack.Context) {
		start := time.Now()

		c.Next()

		fmt.Fprintf(f, "%s %s %s %v %v\n", c.RemoteAddr(), c.Req.Method, c.URL(true), c.Resp.Status(), time.Since(start))
	}
}

// LoggerWithFile returns a baa middleware with file write for log http access
func LoggerWithFile(file string) goblack.HandlerFunc {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("accesslog.LoggerWithFile: open logfile [%s] err: %v", file, err)
	}
	return LoggerWithWriter(f)
}
