#!/bin/bash

#该文件来自AI编写 @newBing （未验证）
#请帮我编写一个bat文件，用于编译go语言代码，接收一个参数可输入目标平台，默认是windows，第二个参数如果输入了则是生成的可执行文件名称，
#文件名称没输入则为app，如果目标平台是windows则添加文件后缀.exe，第三个及以后的参数作为go编译参数，非必填，并且在最后面加上参数：
#-ldflags="-s -w"，最后删除dist目录并重新创建该目录并将app.exe文件和config/目录和public/复制到该目录
#请再帮我写一个用于linux平台的sh文件

# 获取目标平台，默认是linux
target=$1
if [ -z "$target" ]; then
  target=linux
fi

# 获取可执行文件名称，默认是app
name=$2
if [ -z "$name" ]; then
  name=app
fi

# 判断是否需要添加.exe后缀
if [ "$target" == "windows" ]; then
  name=$name.exe
fi

# 获取go编译参数，非必填，并且在最后面加上参数：-ldflags="-s -w"
shift 2 # 移除前两个参数，剩下的作为go编译参数
args="$@ -ldflags=\"-s -w\""

# 删除旧的exe编译文件和dist目录
if [ -f "$name" ]; then
  rm $name
fi

if [ -d "dist" ]; then
  rm -rf dist
fi

# 执行go编译命令，指定目标平台和输出文件名
echo "开始编译go程序为 $target 平台的可执行文件：$name"
GOOS=$target go build -o $name $args

# 判断是否编译成功，如果成功则创建dist目录并复制相关文件到该目录
if [ -f "$name" ]; then
  echo "编译成功！"
  read -p "是否将编译结果打包到dist目录？(y/n)" publish
  if [ "$publish" = "y" ] || [ "$publish" = "Y" ]; then
    echo "编译成功！正在创建dist目录并复制相关文件..."
    mkdir dist
    cp $name dist/
    cp -r config/ dist/config/
    cp -r public/ dist/public/
    echo "完成！请查看dist目录下的内容。"
    fi
else
  echo "编译失败！请检查错误信息。"
fi