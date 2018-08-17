package crm

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type httpCustomer struct {
}

func (p *httpCustomer) onRequest(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/crm/customer/getList") {
		p.onGetList(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/customer/add") {
		p.onAdd(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/customer/edit") {
		p.onEdit(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/customer/delete") {
		p.onDelete(w, r)
	}
}

func (p *httpCustomer) onGetList(w http.ResponseWriter, r *http.Request) {
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

	lst, totalCount, err := instance.db.databaseCustomer.getList(offset, size)
	if err != nil {
		response(w, -1, err.Error())
		return
	}

	var customerList CustomerList
	customerList.CustomerInfos = lst
	customerList.TotalCount = totalCount
	data, err := json.Marshal(customerList)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, string(data))
}

func (p *httpCustomer) onAdd(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	var info CustomerInfo
	info.Id = uuid.Must(uuid.NewV4()).String()
	info.Id = strings.Replace(info.Id, "-", "", -1)
	info.Name = vars.Get("name")
	info.Note = vars.Get("note")
	log.Println(info)
	err := instance.db.databaseCustomer.insert(info)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, "OK")
}

func (p *httpCustomer) onEdit(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	var info CustomerInfo
	info.Id = vars.Get("id")
	info.Name = vars.Get("name")
	info.Note = vars.Get("note")
	log.Println(info)
	err := instance.db.databaseCustomer.update(info)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, "OK")
}

func (p *httpCustomer) onDelete(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	id := vars.Get("id")
	log.Println(id)
	err := instance.db.databaseCustomer.delete(id)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, "OK")
}
