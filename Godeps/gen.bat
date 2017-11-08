set http_proxy=127.0.0.1:1080
set https_proxy=127.0.0.1:1080
git config http.proxy http://127.0.0.1:1080
git config https.proxy https://127.0.0.1:1080

rd /q /s src
set GOPATH=%~dp0
cd %GOPATH%\..
%GOPATH%\..\tool\godep\godep.exe restore
pause