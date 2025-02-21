package router

import (
	"context"
	"fmt"
	"go-helios-da/controller"
	"go-helios-da/logger"
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

	root.Handle(http.MethodGet, "/health", controller.HeliosHealth)

	helios := root.Group("/helios")
	helios.Handle(http.MethodGet, "/hasKey", controller.HeliosHasKey)
	helios.Handle(http.MethodGet, "/getKey", controller.HeliosGetDataByKey)
	helios.Handle(http.MethodGet, "/sugQ", controller.HeliosSugQueryByIndexAndWord)
	helios.Handle(http.MethodGet, "/sugData", controller.HeliosSugDataByIndexAndWord)

	lru := helios.Group("/lru")
	lru.Handle(http.MethodGet, "/keyAll", controller.LRUKeyAll)
	lru.Handle(http.MethodGet, "/get", controller.LRUGetData)
	lru.Handle(http.MethodGet, "/put", controller.LRUPutData)

	// 打印艺术字
	artPrint()

	if resource.RESOURCE_CONF.Port != "" {
		root.Run(":" + resource.RESOURCE_CONF.Port)
		return
	}
	root.Run(":9609")
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
		logger.LOG_USER.Info(u, ip)
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

func artPrint() {
	fmt.Println( "  _    _   ______   _        _____    ____     _____           _____              \n"+
		" | |  | | |  ____| | |      |_   _|  / __ \\   / ____|         |  __ \\      /\\     \n"+
		" | |__| | | |__    | |        | |   | |  | | | (___           | |  | |    /  \\    \n"+
		" |  __  | |  __|   | |        | |   | |  | |  \\___ \\          | |  | |   / /\\ \\   \n"+
		" | |  | | | |____  | |____   _| |_  | |__| |  ____) |    _    | |__| |  / ____ \\  \n"+
		" |_|  |_| |______| |______| |_____|  \\____/  |_____/    (_)   |_____/  /_/    \\_\\ \n")
}
/*
  _    _   ______   _        _____    ____     _____           _____
 | |  | | |  ____| | |      |_   _|  / __ \   / ____|         |  __ \      /\
 | |__| | | |__    | |        | |   | |  | | | (___           | |  | |    /  \
 |  __  | |  __|   | |        | |   | |  | |  \___ \          | |  | |   / /\ \
 | |  | | | |____  | |____   _| |_  | |__| |  ____) |    _    | |__| |  / ____ \
 |_|  |_| |______| |______| |_____|  \____/  |_____/    (_)   |_____/  /_/    \_\

*/