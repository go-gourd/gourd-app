@echo off

@REM 该文件来自AI编写 @newBing
@REM 请帮我编写一个bat文件，用于编译go语言代码，接收一个参数可输入目标平台，默认是windows，第二个参数如果输入了则是生成的可
@REM 执行文件名称，没输入则为app，第三个及以后的参数作为go编译参数，非必填，并且在最后面加上参数：-ldflags="-s -w"，最后删
@REM 除dist目录并重新创建该目录并将app.exe文件和config/目录和public/复制到该目录

rem 设置默认目标平台为windows
set GOOS=windows
rem 设置默认可执行文件名称为app.exe
set NAME=app.exe
rem 获取第一个参数作为目标平台，如果提供的话
if "%1" neq "" set GOOS=%1
rem 获取第二个参数作为可执行文件名称，如果提供的话
if "%2" neq "" set NAME=%2.exe
rem 编译go代码，加上额外的参数和-ldflags="-s -w"
go build -o %NAME% %3 %4 %5 %6 %7 -ldflags="-s -w"
rem 删除dist目录，如果存在的话，并创建一个新的目录
if exist dist rmdir /s /q dist
mkdir dist
rem 复制可执行文件和config/和public/目录到dist/
copy %NAME% dist\
xcopy /e config\ dist\config\
xcopy /e public\ dist\public\