set CURDIR=%~dp0
set BASEDIR=%CURDIR:\src\github.com\fananchong\goredis\=\%
set GOPATH=%BASEDIR%
set GOBIN=%CURDIR%\bin
go install -race ./...
pause