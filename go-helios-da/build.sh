#!/bin/bash

# 进入 Go 工程目录，假设工程目录为 $HOME/go_project
# cd $HOME/go-helios-da

# 构建 Linux 可执行文件
GOOS=linux GOARCH=amd64 go build -o helios-da-linux main.go

# 构建 Windows 可执行文件
GOOS=windows GOARCH=amd64 go build -o helios-da-windows.exe main.go

# 构建 macOS 可执行文件
GOOS=darwin GOARCH=amd64 go build -o helios-da-macos main.go

# 检查 da_conf 文件是否存在
if [ -d "./da_conf" ]; then
    # 打包成 tar.gz
    mkdir log
    tar -czvf helios-da.tar.gz helios-da-linux helios-da-windows.exe helios-da-macos da_conf log
    rm -rf log helios-da-linux helios-da-macos helios-da-windows.exe
    #rm helios-da-windows.exe
    #rm helios-da-macos
    #rm -rf log
else
    echo "da_conf file not found."
    exit 1
fi