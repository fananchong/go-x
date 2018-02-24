set http_proxy=127.0.0.1:1080
set https_proxy=127.0.0.1:1080
git config http.proxy http://127.0.0.1:1080
git config https.proxy https://127.0.0.1:1080

if exist src ( rd /q /s src )
set GOPATH=%~dp0
cd %GOPATH%\..
.\tools\godep\godep.exe restore
cd %GOPATH%
