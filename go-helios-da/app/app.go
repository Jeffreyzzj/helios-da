package app

import (
	"context"
	"fmt"
	"github.com/BurntSushi/toml"
	"go-helios-da/config"
	"go-helios-da/global"
	"go-helios-da/resource"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func InitApp(ctx context.Context) {
	initConf(ctx)

	// 初始化业务日志
	initLog(ctx)
	// 初始化用户日志
	initUserLog(ctx)

	// 初始化脚本
	initShell(ctx)
}

func initConf(ctx context.Context) (err error) {
	var tomlInfo config.TomlConfig
	filePath := global.DA_CONF_PATH
	if _, err = toml.DecodeFile(filePath, &tomlInfo); err != nil {
		fmt.Print(err.Error())
		err = fmt.Errorf("read toml has err %s", err.Error())
		panic(err)
	}
	resource.RESOURCE_CONF = &tomlInfo
	return nil
}

func initShell(ctx context.Context) {
	// 初始化倒排索引
	initTrieTree(ctx)
}

// 初始化倒排索引
func initTrieTree(ctx context.Context) {
	go func() {
		for {
			err := resource.RESOURCE_TRIEROOT.TrieRootInit(ctx)
			if nil != err {
				err = fmt.Errorf("TrieRootInit has err %s \n", err.Error())
				resource.LOGGER.Error("TrieRootInit has err ", zap.Error(err))
			}

			time.Sleep(time.Duration(resource.RESOURCE_CONF.HeliosInitConfig.UpdateTime) * time.Hour)
		}
	}()
}

// 初始化业务日志
func initLog(ctx context.Context) {
	encoder := getEncoder()
	filePath := resource.RESOURCE_CONF.LogInfoPath
	if filePath == "" {
		filePath = global.LOG_INFO_FILE
	}

	fileErrPath := resource.RESOURCE_CONF.LogErrPath
	if fileErrPath == "" {
		fileErrPath = global.LOG_ERR_FILE
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		// 处理文件打开错误
		panic(err.Error())
	}
	errorFile, err := os.OpenFile(fileErrPath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		// 处理文件打开错误
		panic(err.Error())
	}

	infoSync := zapcore.AddSync(file)
	errorSync := zapcore.AddSync(errorFile)

	infoCore := zapcore.NewCore(encoder, infoSync, zapcore.InfoLevel)
	errorCore := zapcore.NewCore(encoder, errorSync, zapcore.ErrorLevel)

	// 使用 zapcore.NewTee 将多个 Core 组合在一起
	teeCore := zapcore.NewTee(infoCore, errorCore)
	//core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	resource.LOGGER = zap.New(teeCore).Sugar()

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