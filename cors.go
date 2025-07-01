package middleware

import (
	"net/http"

	"github.com/varunmbet/goblack"
)

func Cors() goblack.HandlerFunc {
	return func(c *goblack.Context) {
		method := c.Req.Method

		//Access-Control-Expose-Headers <- 允许用户可以读取哪一些 Header
		//Access-Control-Allow-Origin <- 允许哪些域名可以发送请求
		//Access-Control-Allow-Headers <- 允许请求带上哪一些 Header
		//Access-Control-Allow-Methods <- 允许哪些方式发送请求(Get Post Put...等)
		c.SetHeader("Access-Control-Allow-Origin", "*")
		c.SetHeader("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization,Content-Length, x-token")
		c.SetHeader("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.SetHeader("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, Newauthorization, newtoken")
		c.SetHeader("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			//c.AbortWithStatus(http.StatusNoContent)
			c.BreakWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
