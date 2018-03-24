..\..\tools\build\protoc\protoc.exe --go_out=. *.proto
..\..\tools\build\protoc\protoc.exe --js_out=import_style=commonjs,binary:../../tools/h5client/src/app/proto/ *.proto 
pause