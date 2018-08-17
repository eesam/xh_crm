package crm

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type httpDesign struct {
}

func (p *httpDesign) onRequest(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/crm/design/getList") {
		p.onGetList(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/design/add") {
		p.onAdd(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/design/edit") {
		p.onEdit(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/design/delete") {
		p.onDelete(w, r)
	}
}

func (p *httpDesign) onGetList(w http.ResponseWriter, r *http.Request) {
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

	lst, totalCount, err := instance.db.databaseDesign.getList(offset, size)
	if err != nil {
		response(w, -1, err.Error())
		return
	}

	var designList DesignList
	designList.DesignInfos = lst
	designList.TotalCount = totalCount
	data, err := json.Marshal(designList)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, string(data))
}

func (p *httpDesign) onUploadPic(w http.ResponseWriter, r *http.Request) {
	response(w, 0, "OK")
}

func (p *httpDesign) onAdd(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	var info DesignInfo
	info.Id = uuid.Must(uuid.NewV4()).String()
	info.Id = strings.Replace(info.Id, "-", "", -1)
	info.Name = vars.Get("name")
	info.PicUrl = vars.Get("picUrl")
	info.Quantity = 0
	info.Note = vars.Get("note")
	log.Println(info)
	err := instance.db.databaseDesign.insert(info)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, "OK")
}

func (p *httpDesign) onEdit(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	var info DesignInfo
	info.Id = vars.Get("id")
	info.Name = vars.Get("name")
	var err error
	info.Quantity, err = strconv.ParseFloat(vars.Get("quantity"), 10)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	info.Note = vars.Get("note")
	log.Println(info)
	err = instance.db.databaseDesign.update(info)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, "OK")
}

func (p *httpDesign) onDelete(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	id := vars.Get("id")
	log.Println(id)
	err := instance.db.databaseDesign.delete(id)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, "OK")
}
