set CURDIR=%~dp0
set BASEDIR=%CURDIR:\src\github.com\fananchong\go-discovery\=\%
set GOPATH=%BASEDIR%;%CURDIR%\Godeps
set GOBIN=%CURDIR%\bin
go install -race ./example

pause
