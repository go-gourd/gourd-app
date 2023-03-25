@echo off

@REM 该文件来自AI编写 @newBing
@REM 请帮我编写一个bat文件，用于编译go语言代码，接收一个参数可输入目标平台，默认是windows，第二个参数如果输入了则是生成的可执行文件名称，
@REM 文件名称没输入则为app，如果目标平台是windows则添加文件后缀.exe，第三个及以后的参数作为go编译参数，非必填，并且在最后面加上参数：
@REM -ldflags="-s -w"，最后删除dist目录并重新创建该目录并将app.exe文件和config/目录和public/复制到该目录

:: 获取目标平台，默认是windows
set target=%1
if "%target%" == "" set target=windows

:: 获取可执行文件名称，默认是app
set name=%2
if "%name%" == "" set name=app

:: 判断是否需要添加.exe后缀
if "%target%" == "windows" set name=%name%.exe

:: 获取go编译参数，非必填，并且在最后面加上参数：-ldflags="-s -w"
set args=%3 %4 %5 %6 %7 %8 %9 -ldflags="-s -w"

:: 删除旧的编译文件
if exist %name% del %name%

:: 执行go编译命令，指定目标平台和输出文件名
echo 开始编译为%target%平台的可执行文件：%name%
set GOOS=%target%
go build -o %name% %args%

:: 判断是否编译成功，如果成功则创建dist目录并复制相关文件到该目录
if exist %name%  (
    echo 编译成功！
) else (
    echo 编译失败！请检查错误信息。
    exit
)

:: 询问用户是否创建发布
set /p choice=是否将编译结果打包到dist目录？(y/n)

if "%choice%" == "y" (
    rd /s /q dist
    md dist
    copy /y %name% dist\
    xcopy /e /y config\ dist\config\
    xcopy /e /y public\ dist\public\
    del %name%
    echo 完成！请查看dist目录下的内容。
)