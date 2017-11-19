
set ETCD_HOSTS=192.168.1.4:12379,192.168.1.4:22379,192.168.1.4:32379

start base.exe --etcdHosts=%ETCD_HOSTS% --log_dir=./log -stderrthreshold 0
start gateway.exe --etcdHosts=%ETCD_HOSTS% --log_dir=./log -stderrthreshold 0
start room.exe --etcdHosts=%ETCD_HOSTS% --log_dir=./log -stderrthreshold 0
start login.exe --etcdHosts=%ETCD_HOSTS% --etcdWatchNodeTypes=2,4 --log_dir=./log -stderrthreshold 0
