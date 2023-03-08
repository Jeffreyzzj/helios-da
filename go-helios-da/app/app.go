package app

import (
	"context"
	"go-helios-da/global"
	"go-helios-da/resource"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func InitApp(ctx context.Context) {
	// 初始化业务日志
	initLog(ctx)
	// 初始化用户日志
	initUserLog(ctx)

}

// 初始化业务日志
func initLog(ctx context.Context) {
	encoder := getEncoder()
	file, _ := os.OpenFile(global.LOG_INFO_FILE, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	sync := getWriteSync(file)
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	resource.LOGGER = zap.New(core).Sugar()

}

// 初始化用户日志
func initUserLog(ctx context.Context) {
	encoder := getEncoder()
	file, _ := os.OpenFile(global.LOG_USER_INFO_FILE, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	sync := getWriteSync(file)
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	resource.LOGGER_USER = zap.New(core).Sugar()

}

// 负责设置 encoding 的日志格式
func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

// 负责日志写入的位置
func getWriteSync(file *os.File) zapcore.WriteSyncer {
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
