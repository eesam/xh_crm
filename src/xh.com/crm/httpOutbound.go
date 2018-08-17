package crm

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type httpOutbound struct {
}

func (p *httpOutbound) onRequest(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/crm/outbound/getList") {
		p.onGetList(w, r)
	}
}

func (p *httpOutbound) onGetList(w http.ResponseWriter, r *http.Request) {
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
	lst, totalCount, err := instance.db.databaseOutbound.getList(offset, size)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	var outboundList OutboundList
	outboundList.OutboundInfos = lst
	outboundList.TotalCount = totalCount
	data, err := json.Marshal(outboundList)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, string(data))
}
