package common

import "net/http"

type WebService struct {
	server      *http.Server
	serverMux   *http.ServeMux
	termination bool
}

func NewWebService() *WebService {
	return &WebService{}
}

func (this *WebService) ListenAndServe(addr string) {
	xlog.Infoln("start listen", addr)
	this.termination = false
	for !this.termination {
		serverMux := http.NewServeMux()
		this.server = &http.Server{Addr: addr, Handler: serverMux}
		this.serverMux = serverMux
		err := this.server.ListenAndServe()
		if err != nil {
			xlog.Errorln("[web]", err)
			this.server.Close()
		}
	}
}

func (this *WebService) Close() {
	this.termination = true
	if this.server != nil {
		this.server.Close()
		this.server = nil
		this.serverMux = nil
	}
}

func (this *WebService) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	if this.serverMux != nil {
		this.serverMux.HandleFunc(pattern, handler)
	}
}
