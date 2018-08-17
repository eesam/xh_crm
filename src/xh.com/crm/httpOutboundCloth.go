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

type httpOutboundCloth struct {
}

func (p *httpOutboundCloth) onRequest(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/crm/outboundCloth/getList") {
		p.onGetList(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/outboundCloth/add") {
		p.onAdd(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/outboundCloth/edit") {
		p.onEdit(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/crm/outboundCloth/delete") {
		p.onDelete(w, r)
	}
}

func (p *httpOutboundCloth) onGetList(w http.ResponseWriter, r *http.Request) {
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

	lst, totalCount, err := instance.db.databaseOutboundCloth.getList(offset, size)
	if err != nil {
		response(w, -1, err.Error())
		return
	}

	var outboundClothList OutboundClothList
	outboundClothList.OutboundClothInfos = lst
	outboundClothList.TotalCount = totalCount
	data, err := json.Marshal(outboundClothList)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, string(data))
}

func (p *httpOutboundCloth) onAdd(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	log.Printf("%+v\n", vars)

	outboundCloth := struct {
		CustomerId string `json:"customerId"`
		Note       string `json:"note"`
		Cloths     []struct {
			InboundClothId   string  `json:"inboundClothId"`
			OutboundPrice    float64 `json:"outboundPrice"`
			OutboundQuantity float64 `json:"outboundQuantity"`
		} `json:"cloths"`
	}{}
	outboundCloths := vars["outboundCloths"]

	outboundId := uuid.Must(uuid.NewV4()).String()
	outboundId = strings.Replace(outboundId, "-", "", -1)
	for _, str := range outboundCloths {
		err := json.Unmarshal([]byte(str), &outboundCloth)
		if err != nil {
			response(w, -1, err.Error())
			return
		}

		for _, cloth := range outboundCloth.Cloths {
			var info OutboundClothInfo
			info.Id = uuid.Must(uuid.NewV4()).String()
			info.Id = strings.Replace(info.Id, "-", "", -1)
			info.CustomerId = outboundCloth.CustomerId
			info.Time = fmt.Sprintf("%s", time.Now().Format("2006-01-02 15:04:05"))
			info.Note = outboundCloth.Note

			info.InboundClothId = cloth.InboundClothId
			info.Quantity = cloth.OutboundQuantity
			info.Price = cloth.OutboundPrice

			err = instance.db.databaseOutboundCloth.insert(info, outboundId)
			if err != nil {
				response(w, -1, err.Error())
				return
			}
		}
	}
	response(w, 0, "OK")
}

func (p *httpOutboundCloth) onEdit(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	var info OutboundClothInfo
	info.Id = vars.Get("id")
	info.DesignId = vars.Get("designId")
	info.ColorId = vars.Get("colorId")
	info.CustomerId = vars.Get("customerId")
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
	err = instance.db.databaseOutboundCloth.update(info)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, "OK")
}

func (p *httpOutboundCloth) onDelete(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	id := vars.Get("id")
	log.Println(id)
	err := instance.db.databaseOutboundCloth.delete(id)
	if err != nil {
		response(w, -1, err.Error())
		return
	}
	response(w, 0, "OK")
}
