set MYROOT=%~dp0

cd %MYROOT%
cd common_services
call build.bat

cd %MYROOT%
cd example1_iogame
call build.bat

cd %MYROOT%
cd test\test_packet_loss
call build.bat

cd %MYROOT%

pause