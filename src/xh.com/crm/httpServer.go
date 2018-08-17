package crm

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type httpServer struct {
	httpColor
	httpCustomer
	httpSupplier
	httpDesign
	httpInboundCloth
	httpOutboundCloth
	httpInbound
	httpOutbound
	httpScheduling
	httpDesignColor
}

func newHttpServer() *httpServer {
	p := new(httpServer)
	return p
}

type httpResp struct {
	Code int
	Msg  string
}

func (p *httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, HEAD")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Max-Age", "30")

	if strings.HasPrefix(r.URL.Path, "/crm/color") {
		p.httpColor.onRequest(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/customer/") {
		p.httpCustomer.onRequest(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/supplier/") {
		p.httpSupplier.onRequest(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/design/") {
		p.httpDesign.onRequest(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/inboundCloth/") {
		p.httpInboundCloth.onRequest(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/outboundCloth/") {
		p.httpOutboundCloth.onRequest(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/inbound/") {
		p.httpInbound.onRequest(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/outbound/") {
		p.httpOutbound.onRequest(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/scheduling/") {
		p.httpScheduling.onRequest(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/designColor/") {
		p.httpDesignColor.onRequest(w, r)
	}
}

func (p *httpServer) run() {
	fs := http.FileServer(http.Dir("/root/xh_crm/web"))
	http.Handle("/", fs)
	http.HandleFunc("/crm/", p.ServeHTTP)
	log.Fatal(http.ListenAndServe("0.0.0.0:8081", nil))
}

func response(w http.ResponseWriter, code int, msg string) {
	resp := httpResp{}
	resp.Code = code
	resp.Msg = msg
	data, err := json.MarshalIndent(resp, "", "  ")
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	} else {
		w.WriteHeader(500)
		fmt.Fprintf(w, "%#v", err)
	}
}
