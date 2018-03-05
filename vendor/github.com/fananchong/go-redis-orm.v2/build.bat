set CURDIR=%~dp0
set BASEDIR=%CURDIR:\src\github.com\fananchong\go-redis-orm.v2\=\%
set GOPATH=%BASEDIR%

set GOBIN=%CURDIR%\release
go install ./tools/redis2go/...

cd .\example
call g.bat
cd %CURDIR%
set GOBIN=%CURDIR%\example
go install ./example/...

pause