set CURDIR=%~dp0
set BASEDIR=%CURDIR:\src\github.com\fananchong\go-x\test\test_packet_loss=\%
set GOPATH=%BASEDIR%
set GOBIN=%CURDIR%\bin
go install ./...