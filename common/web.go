package common

import (
	"net/http"
	"time"
)

type WebService struct {
	server      *http.Server
	serverMux   *http.ServeMux
	termination bool
}

func NewWebService() *WebService {
	return &WebService{serverMux: http.NewServeMux()}
}

func (this *WebService) ListenAndServe(addr string) {
	xlog.Infoln("start listen", addr)
	this.termination = false
	for !this.termination {
		this.server = &http.Server{Addr: addr, Handler: this.serverMux}
		err := this.server.ListenAndServe()
		if err != nil {
			xlog.Errorln("[web]", err)
			this.server.Close()
			time.Sleep(5 * time.Second)
		}
	}
}

func (this *WebService) Close() {
	this.termination = true
	if this.server != nil {
		this.server.Close()
		this.server = nil
	}
}

func (this *WebService) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	if this.serverMux == nil {
		this.serverMux = http.NewServeMux()
	}
	this.serverMux.HandleFunc(pattern, handler)
}
