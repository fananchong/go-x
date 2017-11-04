rd /q /s src
set GOPATH=%~dp0
cd %GOPATH%\..
%GOPATH%\..\tool\godep\godep.exe restore
pause