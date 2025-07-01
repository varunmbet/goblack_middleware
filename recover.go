package middleware

import (
	"fmt"

	"runtime"

	"github.com/varunmbet/goblack"
)

func Recovery() goblack.HandlerFunc {
	return func(c *goblack.Context) {
		defer func() {
			if err := recover(); err != nil {
				trace := make([]byte, 1<<16)
				n := runtime.Stack(trace, true)
				c.Error(fmt.Errorf("panic recover\n %v\n stack trace %d bytes\n %s", err, n, trace[:n]))
			}
		}()

		c.Next()
	}
}
