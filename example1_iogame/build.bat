set CURDIR=%~dp0
set BASEDIR=%CURDIR:\src\github.com\fananchong\go-x\example1_iogame\=\%
set GOPATH=%BASEDIR%;%CURDIR%
set GOBIN=%CURDIR%\bin
go install -race ./...
