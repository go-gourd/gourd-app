
# ���ļ����ڱ���go���Դ��룬����AI���� @phind.com
# ����ұ�дһ��powershell�ļ������ڱ���go���Դ��룬����һ������������Ŀ��ƽ̨��Ĭ����windows���ڶ�
# ����������������������ɵĿ�ִ���ļ����ƣ��ļ�����û������Ϊapp�����Ŀ��ƽ̨��windows������ļ���
# ׺.exe�����������Ժ�Ĳ�����Ϊgo����������Ǳ���������������ϲ�����-ldflags="-s -w",�ڱ���
# ǰɾ��֮ǰ���ɵĶ������ļ�����ж��Ƿ����ɹ����������ʾ

param(
    [Parameter(Position=0)]
    [string]$targetPlatform = "windows",

    [Parameter(Position=1)]
    [string]$outputName = "app",

    [Parameter(Position=2, ValueFromRemainingArguments=$true)]
    [string[]]$goBuildParams = @()
)

# ���Ŀ��ƽ̨��windows�����.exe��׺
if($targetPlatform -eq "windows") {
    $outputName += ".exe"
}

# ����GOOS��GOARCH��������
$env:GOOS = $targetPlatform
$env:GOARCH = "amd64"

# ����go build����
$command = "go build -o $outputName"
$command += " -ldflags=`"-s -w`""
if($goBuildParams.Length -ne 0) {
    $command += " " + [string]::Join(" ", $goBuildParams)
}

# �������ļ��Ѵ��ڣ�ɾ����
if(Test-Path $outputName) {
    Remove-Item $outputName
}

Write-Output $command
# ��������
Invoke-Expression $command

# �������Ƿ�ɹ�
if(Test-Path $outputName) {
    Write-Output "����ɹ������ɵ��ļ���Ϊ��$outputName"
} else {
    Write-Output "����ʧ��"
}
