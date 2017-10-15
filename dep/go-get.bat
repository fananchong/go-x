
set http_proxy=127.0.0.1:1080
set https_proxy=127.0.0.1:1080

git config http.proxy http://127.0.0.1:1080
git config https.proxy https://127.0.0.1:1080

set GOPATH=%~dp0
echo %GOPATH%

go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/glog
go get -u github.com/fananchong/gonet
go get -u github.com/fananchong/gonet_example
go get -u github.com/fananchong/go-proto-helper

go get -u github.com/gonuts/commander
go get -u github.com/StackExchange/wmi
go get -u golang.org/x/sys/windows
go get -u github.com/shirou/w32
go get -u github.com/shirou/gopsutil
go get -u github.com/fananchong/gochart
go get -u github.com/fananchong/gochart_example


pause