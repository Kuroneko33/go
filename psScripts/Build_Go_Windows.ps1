Param 
(
    [string]$path,
    [string]$file
)

$ThisLocation = Get-Location
Write-Host $ThisLocation

$Env:GOOS = "windows"
$Env:GOARCH = "amd64"
$Env:GO111MODULE = "off"
#$Env:GOPATH = $ThisLocation
$Env:GOPATH = "C:\Users\55541\source\goProjects"
go env GOOS GOARCH GO111MODULE GOPATH

if(!$path)
{
    $path = "./"
}

if(!$file)
{
    $file = $path + "main.exe"
}

Write-Host $file
go build -o $file $path