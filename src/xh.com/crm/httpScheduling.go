package crm

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type httpScheduling struct {
}

func (p *httpScheduling) onRequest(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/crm/scheduling/getList") {
		p.onGetList(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/scheduling/add") {
		p.onAdd(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/scheduling/edit") {
		p.onEdit(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/scheduling/delete") {
		p.onDelete(w, r)
	}
}

func (p *httpScheduling) onGetList(w http.ResponseWriter, r *http.Request) {
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

	lst, totalCount, err := instance.db.databaseScheduling.getList(offset, size)
	if err != nil {
		response(w, -1, err.Error())
		return
	}

	var schedulingList SchedulingList
	schedulingList.SchedulingInfos = lst
	schedulingList.TotalCount = totalCount
	data, err := json.Marshal(schedulingList)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, string(data))
}

func (p *httpScheduling) onAdd(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	var info SchedulingInfo
	info.Id = uuid.Must(uuid.NewV4()).String()
	info.Id = strings.Replace(info.Id, "-", "", -1)
	info.DesignId = vars.Get("designId")
	info.ColorId = vars.Get("colorId")
	info.SupplierId = vars.Get("supplierId")
	var err error
	info.Quantity, err = strconv.ParseFloat(vars.Get("quantity"), 10)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	info.Price, err = strconv.ParseFloat(vars.Get("price"), 10)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	info.Time = vars.Get("time")
	if info.Time == "" {
		info.Time = fmt.Sprintf("%s", time.Now().Format("2006-01-02 15:04:05"))
	}
	info.Note = vars.Get("note")
	log.Println(info)
	err = instance.db.databaseScheduling.insert(info)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, "OK")
}

func (p *httpScheduling) onEdit(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	var info SchedulingInfo
	info.Id = vars.Get("id")
	info.DesignId = vars.Get("designId")
	info.ColorId = vars.Get("colorId")
	info.SupplierId = vars.Get("supplierId")
	var err error
	info.Quantity, err = strconv.ParseFloat(vars.Get("quantity"), 10)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	info.Price, err = strconv.ParseFloat(vars.Get("price"), 10)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	info.Time = vars.Get("time")
	if info.Time == "" {
		info.Time = fmt.Sprintf("%s", time.Now().Format("2006-01-02 15:04:05"))
	}
	info.Note = vars.Get("note")
	log.Println(info)
	err = instance.db.databaseScheduling.update(info)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, "OK")
}

func (p *httpScheduling) onDelete(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	id := vars.Get("id")
	log.Println(id)
	err := instance.db.databaseScheduling.delete(id)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, "OK")
}
