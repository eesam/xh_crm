package crm

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type httpInbound struct {
}

func (p *httpInbound) onRequest(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/crm/inbound/getList") {
		p.onGetList(w, r)
	}
}

func (p *httpInbound) onGetList(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	offset, err := strconv.Atoi(vars.Get("offset"))
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	size, err := strconv.Atoi(vars.Get("size"))
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	lst, totalCount, err := instance.db.databaseInbound.getList(offset, size)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	var inboundList InboundList
	inboundList.InboundInfos = lst
	inboundList.TotalCount = totalCount
	data, err := json.Marshal(inboundList)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, string(data))
}
