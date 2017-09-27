package http

import (
	"github.com/open-falcon/falcon-plus/modules/agent/g"
	"net/http"
)

func bindRoutes(route string, handler func(http.ResponseWriter, *http.Request)) {

	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		if g.IsTrustable(r.RemoteAddr) {
			handler(w, r)
		}else{
			w.Write([]byte("no privilege"))
		}
	})

}		