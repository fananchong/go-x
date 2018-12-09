set CURDIR=%~dp0
set BASEDIR=%CURDIR%\..\..\..\..\..\
set GOPATH=%BASEDIR%;%CURDIR%
set GOBIN=%CURDIR%\bin
go install -race ./...
