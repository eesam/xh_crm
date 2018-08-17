package crm

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type httpSupplier struct {
}

func (p *httpSupplier) onRequest(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/crm/supplier/getList") {
		p.onGetList(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/supplier/add") {
		p.onAdd(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/supplier/edit") {
		p.onEdit(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/supplier/delete") {
		p.onDelete(w, r)
	}
}

func (p *httpSupplier) onGetList(w http.ResponseWriter, r *http.Request) {
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

	lst, totalCount, err := instance.db.databaseSupplier.getList(offset, size)
	if err != nil {
		response(w, -1, err.Error())
		return
	}

	var supplierList SupplierList
	supplierList.SupplierInfos = lst
	supplierList.TotalCount = totalCount
	data, err := json.Marshal(supplierList)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, string(data))
}

func (p *httpSupplier) onAdd(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	var info SupplierInfo
	info.Id = uuid.Must(uuid.NewV4()).String()
	info.Id = strings.Replace(info.Id, "-", "", -1)
	info.Name = vars.Get("name")
	info.Note = vars.Get("note")
	log.Println(info)
	err := instance.db.databaseSupplier.insert(info)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, "OK")
}

func (p *httpSupplier) onEdit(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	var info SupplierInfo
	info.Id = vars.Get("id")
	info.Name = vars.Get("name")
	info.Note = vars.Get("note")
	log.Println(info)
	err := instance.db.databaseSupplier.update(info)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, "OK")
}

func (p *httpSupplier) onDelete(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	id := vars.Get("id")
	log.Println(id)
	err := instance.db.databaseSupplier.delete(id)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, "OK")
}
