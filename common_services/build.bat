set CURDIR=%~dp0
set BASEDIR=%CURDIR:\src\github.com\fananchong\go-x\common_services\=\%
set GOPATH=%BASEDIR%;%CURDIR%
set GOBIN=%CURDIR%\bin
go install -race ./...
