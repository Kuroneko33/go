Param 
(
    [string]$path,
    [string]$file
)

$ThisLocation = Get-Location
Write-Host $ThisLocation

$Env:GOOS = "linux"
$Env:GOARCH = "amd64"
$Env:GO111MODULE = "off"
$Env:GOPATH = $ThisLocation
go env GOOS GOARCH GO111MODULE GOPATH

if(!$path)
{
    $path = "./"
}

if(!$file)
{
    $file = $path + "main.linux"
}

Write-Host $file
go build -o $file $path