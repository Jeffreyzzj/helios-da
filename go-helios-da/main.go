package main

import (
	"context"
	"go-helios-da/app"
	"go-helios-da/router"
)

func main() {
	ctx := context.Background()
	// 初始化
	app.InitApp(ctx)

	// 处理请求
	router.Router(ctx)
}
