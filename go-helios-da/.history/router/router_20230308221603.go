package router

import (
	"context"
	"fmt"
	"go-helios-da/resource"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Router(ctx context.Context) {
	root := gin.Default()
	root.Use(Core())

	testGroup := root.Group("/test")
	testGroup.Handle(http.MethodGet, "health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	root.Run(":8911")
}

func Core() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token")
		c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type")
		c.Header("Access-Control-Allow-Credentials", "True")
		//放行索引options
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		//处理请求
		c.Next()
	}
}

// 自定义 Logger() 中间件
func HttpLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now().Format("2006-01-02 15:03:04")
		ip := ClientPublicIP(c.Request)

		path := c.FullPath()
		query := c.DefaultQuery("q", "") //带默认值

		u := fmt.Sprintf(`|{"t":"%s", "ip":"%s", "fun":"%s", "q":"%s"}|`,
			t, ip, path, query)
		resource.LOGGER_USER.Info(u, ip)
		c.Next()
	}
}

func ClientPublicIP(r *http.Request) string {
	var ip string

	if r.Header.Get("X-Forwarded-For") != "" {
		return r.Header.Get("X-Forwarded-For")
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" && !HasLocalIPddr(ip) {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("x-client-ip"))
	if ip != "" && !HasLocalIPddr(ip) {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

func HasLocalIPddr(ip string) bool {
	return true
}
