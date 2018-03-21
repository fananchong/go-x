set CURPATH=%~dp0\src
echo %CURPATH%

if not exist %CURPATH%\..\userdata (mkdir %CURPATH%\..\userdata)

"C:\Program Files (x86)\Google\Chrome\Application\chrome.exe" --allow-file-access-from-files --disable-web-security --user-data-dir=%CURPATH%\..\userdata --app=file:///%CURPATH%/index.html
