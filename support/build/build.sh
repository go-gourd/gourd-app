#!/bin/bash

#该文件来自AI编写 @newBing （未验证）
#请帮我编写一个bat文件，用于编译go语言代码，接收一个参数可输入目标平台，默认是windows，第二个参数如果输入了则是生成的可
#执行文件名称，没输入则为app，第三个及以后的参数作为go编译参数，非必填，并且在最后面加上参数：-ldflags="-s -w"，最后删
#除dist目录并重新创建该目录并将app.exe文件和config/目录和public/复制到该目录
#请再帮我写一个用于linux平台的sh文件

# 设置默认目标平台为linux
GOOS=linux
# 设置默认可执行文件名称为app
NAME=app
# 获取第一个参数作为目标平台，如果提供的话
if [ -n "$1" ]; then GOOS=$1; fi
# 获取第二个参数作为可执行文件名称，如果提供的话
if [ -n "$2" ]; then NAME=$2; fi
# 编译go代码，加上额外的参数和-ldflags="-s -w"
go build -o $NAME ${@:3} -ldflags="-s -w"
# 删除dist目录，如果存在的话，并创建一个新的目录
rm -rf dist && mkdir dist
# 复制可执行文件和config/和public/目录到dist/
cp $NAME dist/
cp -r config/ dist/config/
cp -r public/ dist/public/