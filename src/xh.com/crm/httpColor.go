package crm

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type httpColor struct {
}

func (p *httpColor) onRequest(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/crm/color/getList") {
		p.onGetList(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/color/add") {
		p.onAdd(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/color/edit") {
		p.onEdit(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/color/delete") {
		p.onDelete(w, r)
	}
}

func (p *httpColor) onGetList(w http.ResponseWriter, r *http.Request) {
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

	lst, totalCount, err := instance.db.databaseColor.getList(offset, size)
	if err != nil {
		response(w, -1, err.Error())
		return
	}

	var colorList ColorList
	colorList.ColorInfos = lst
	colorList.TotalCount = totalCount
	data, err := json.Marshal(colorList)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, string(data))
}

func (p *httpColor) onAdd(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	var info ColorInfo
	info.Id = uuid.Must(uuid.NewV4()).String()
	info.Id = strings.Replace(info.Id, "-", "", -1)
	info.Name = vars.Get("name")
	info.Note = vars.Get("note")
	log.Println(info)
	err := instance.db.databaseColor.insert(info)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, "OK")
}

func (p *httpColor) onEdit(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	var info ColorInfo
	info.Id = vars.Get("id")
	info.Name = vars.Get("name")
	info.Note = vars.Get("note")
	log.Println(info)
	err := instance.db.databaseColor.update(info)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, "OK")
}

func (p *httpColor) onDelete(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	id := vars.Get("id")
	log.Println(id)
	err := instance.db.databaseColor.delete(id)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, "OK")
}
