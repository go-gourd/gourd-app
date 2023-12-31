@echo off

@REM 此文件用于编译go语言代码，来自AI生成 @phind.com
@REM 请帮我编写一个bat文件，用于编译go语言代码，接收一个参数可输入目标平台，默认是windows，第二个参数如果输入了则是生成的可执行文件名称，
@REM 文件名称没输入则为app，如果目标平台是windows则添加文件后缀.exe，第三个及以后的参数作为go编译参数，非必填，并且在最后面加上参数：
@REM -ldflags="-s -w"

:: 获取目标平台，默认为windows
set TARGET_PLATFORM=%1
if "%TARGET_PLATFORM%"=="" set TARGET_PLATFORM=windows

:: 获取生成的可执行文件名称，默认为app
set OUTPUT_NAME=%2
if "%OUTPUT_NAME%"=="" set OUTPUT_NAME=app

:: 如果目标平台是windows，则添加文件后缀.exe
if "%TARGET_PLATFORM%"=="windows" set OUTPUT_NAME=%OUTPUT_NAME%.exe

:: 删除之前生成的二进制文件
if exist %OUTPUT_NAME% del /f /q %OUTPUT_NAME%

:: 获取go编译参数
set GO_BUILD_ARGS=%3

:: 编译go代码
go build -o %OUTPUT_NAME% -ldflags="-s -w" %GO_BUILD_ARGS%

:: 判断是否编译成功
if not exist %OUTPUT_NAME% (
   echo Compile failed!
   exit /b 1
)

echo Compile successful!
exit /b 0
