set MYROOT=%~dp0

cd example1_iogame
call build.bat

cd %MYROOT%
cd test\test_packet_loss
call build.bat

pause