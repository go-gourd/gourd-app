
# 此文件用于编译go语言代码，来自AI生成 @phind.com
# 请帮我编写一个powershell文件，用于编译go语言代码，接收一个参数可输入目标平台，默认是windows，第二
# 个参数如果输入了则是生成的可执行文件名称，文件名称没输入则为app，如果目标平台是windows则添加文件后
# 缀.exe，第三个及以后的参数作为go编译参数，非必填，并且在最后面加上参数：-ldflags="-s -w",在编译
# 前删除之前生成的二进制文件最后判断是否编译成功，并输出提示

param(
    [Parameter(Position=0)]
    [string]$targetPlatform = "windows",

    [Parameter(Position=1)]
    [string]$outputName = "app",

    [Parameter(Position=2, ValueFromRemainingArguments=$true)]
    [string[]]$goBuildParams = @()
)

# 如果目标平台是windows，添加.exe后缀
if($targetPlatform -eq "windows") {
    $outputName += ".exe"
}

# 设置GOOS和GOARCH环境变量
$env:GOOS = $targetPlatform
$env:GOARCH = "amd64"

# 构建go build命令
$command = "go build -o $outputName"
$command += " -ldflags=`"-s -w`""
if($goBuildParams.Length -ne 0) {
    $command += " " + [string]::Join(" ", $goBuildParams)
}

# 如果输出文件已存在，删除它
if(Test-Path $outputName) {
    Remove-Item $outputName
}

Write-Output $command
# 运行命令
Invoke-Expression $command

# 检查编译是否成功
if(Test-Path $outputName) {
    Write-Output "编译成功，生成的文件名为：$outputName"
} else {
    Write-Output "编译失败"
}
