set CURDIR=%~dp0
set BASEDIR=%CURDIR:\src\github.com\fananchong\go-x\=\%
set GOPATH=%BASEDIR%;%CURDIR%\Godeps
set GOBIN=%CURDIR%\bin
go install ./server/...
go install ./test/...
pause