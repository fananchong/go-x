
set http_proxy=127.0.0.1:1080
set https_proxy=127.0.0.1:1080

git config http.proxy http://127.0.0.1:1080
git config https.proxy https://127.0.0.1:1080

set GOPATH=%~dp0
echo %GOPATH%

go get -u -d github.com/fananchong/glog
go get -u -d github.com/fananchong/gotcp
go get -u -d github.com/fananchong/multiconfig
go get -u -d github.com/fananchong/gomap
go get -u -d github.com/fananchong/go-redis-orm.v2
go get -u -d github.com/fananchong/goredis
go get -u -d github.com/fananchong/go-discovery
go get -u -d github.com/fananchong/gochart
go get -u -d github.com/golang/protobuf/proto
go get -u -d github.com/gogo/protobuf/proto
go get -u -d github.com/garyburd/redigo/redis
go get -u -d github.com/gomodule/redigo/redis
go get -u -d github.com/satori/go.uuid
go get -u -d github.com/FZambia/go-sentinel
go get -u -d github.com/bitly/go-simplejson
go get -u -d github.com/coreos/etcd
go get -u -d github.com/mna/redisc
go get -u -d golang.org/x/net/...
go get -u -d golang.org/x/text/...
go get -u -d google.golang.org/genproto/...

pause