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

type httpInboundCloth struct {
}

func (p *httpInboundCloth) onRequest(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/crm/inboundCloth/getList") {
		p.onGetList(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/inboundCloth/add") {
		p.onAdd(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/inboundCloth/edit") {
		p.onEdit(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/inboundCloth/delete") {
		p.onDelete(w, r)
	}
}

func (p *httpInboundCloth) onGetList(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	designId := vars.Get("designId")
	colorId := vars.Get("colorId")
	lst, totalCount, err := instance.db.databaseInboundCloth.getList(designId, colorId)
	if err != nil {
		response(w, -1, err.Error())
		return
	}

	var inboundClothList InboundClothList
	inboundClothList.InboundClothInfos = lst
	inboundClothList.TotalCount = totalCount
	data, err := json.Marshal(inboundClothList)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, string(data))
}

func (p *httpInboundCloth) onAdd(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	log.Printf("%+v\n", vars)

	inboundCloth := struct {
		DesignId   string    `json:"designId"`
		ColorId    string    `json:"colorId"`
		SupplierId string    `json:"supplierId"`
		Quantitys  []float64 `json:"quantitys"`
		Price      float64   `json:"price"`
		Note       string    `json:"note"`
	}{}
	inboundCloths := vars["inboundCloths"]

	inboundId := uuid.Must(uuid.NewV4()).String()
	inboundId = strings.Replace(inboundId, "-", "", -1)
	for _, str := range inboundCloths {
		err := json.Unmarshal([]byte(str), &inboundCloth)
		if err != nil {
			response(w, -1, err.Error())
			return
		}

		for _, quantity := range inboundCloth.Quantitys {
			var info InboundClothInfo
			info.Id = uuid.Must(uuid.NewV4()).String()
			info.Id = strings.Replace(info.Id, "-", "", -1)
			info.DesignId = inboundCloth.DesignId
			info.ColorId = inboundCloth.ColorId
			info.SupplierId = inboundCloth.SupplierId
			info.Quantity = quantity
			info.RemainQuantity = quantity
			info.Price = inboundCloth.Price
			info.Note = inboundCloth.Note
			info.Time = fmt.Sprintf("%s", time.Now().Format("2006-01-02 15:04:05"))
			err = instance.db.databaseInboundCloth.insert(info, inboundId)
			if err != nil {
				response(w, -1, err.Error())
				return
			}
		}
	}
	response(w, 0, "OK")
}

func (p *httpInboundCloth) onEdit(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	var info InboundClothInfo
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
	err = instance.db.databaseInboundCloth.update(info)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, "OK")
}

func (p *httpInboundCloth) onDelete(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	id := vars.Get("id")
	log.Println(id)
	err := instance.db.databaseInboundCloth.delete(id)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, "OK")
}
