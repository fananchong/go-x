set CURDIR=%~dp0
set BASEDIR=%CURDIR:\src\github.com\fananchong\gomap\=\%
set GOPATH=%BASEDIR%
set GOBIN=%CURDIR%\bin
go install ./...
pause