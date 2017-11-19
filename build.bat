set CURDIR=%~dp0
set BASEDIR=%CURDIR:\src\github.com\fananchong\go-x\=\%
set GOPATH=%BASEDIR%;%CURDIR%\Godeps
set GOBIN=%CURDIR%\bin
go install -race ./server/...
go install -race ./test/...
pause