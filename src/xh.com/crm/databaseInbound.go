package crm

import (
	"database/sql"
	"errors"
	"log"
)

type inboundInfo struct {
	Id             string
	InboundClothId string
	DesignId       string
	DesignName     string
	ColorId        string
	ColorName      string
	SupplierId     string
	SupplierName   string
	Quantity       float64
	RemainQuantity float64
	Price          float64
	Time           string
	Note           string
}

type Inbound struct {
	Id             string  `json:"id"`
	DesignId       string  `json:"designId"`
	DesignName     string  `json:"designName"`
	ColorId        string  `json:"colorId"`
	ColorName      string  `json:"colorName"`
	TotalQuantity  float64 `json:"totalQuantity"`
	RemainQuantity float64 `json:"remainQuantity"`
	Price          float64 `json:"price"`
	Time           string  `json:"time"`
	Note           string  `json:"note"`
}

type InboundInfo struct {
	Id       string    `json:"id"`
	Inbounds []Inbound `json:"inbounds"`
}

type InboundList struct {
	TotalCount   int           `json:"totalCount"`
	InboundInfos []InboundInfo `json:"inboundInfos"`
}

type databaseInbound struct {
	db *sql.DB
}

func (p *databaseInbound) getList(offset int, size int) ([]InboundInfo, int, error) {
	if p.db == nil {
		return nil, 0, errors.New("")
	}

	var rows *sql.Rows
	var err error

	////////////////////////////////////////////////////////
	// totalCount
	query1 := "SELECT count(DISTINCT `inbound_id`) FROM `xh_crm`.`xh_inbound`"
	log.Println(query1)
	rows, err = p.db.Query(query1)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	totalCount := 0
	if rows.Next() {
		err = rows.Scan(&totalCount)
		if err != nil {
			log.Println(err)
			return nil, 0, err
		}
		log.Println("totalCount", totalCount)
	}
	////////////////////////////////////////////////////////

	query2 := "SELECT `xh_inbound`.`inbound_id`" +
		",`xh_inbound_cloth`.`inbound_cloth_id`" +
		",`xh_inbound_cloth`.`inbound_cloth_quantity`" +
		",`xh_inbound_cloth`.`remain_quantity`" +
		",`xh_inbound_cloth`.`inbound_cloth_price`" +
		",`xh_inbound_cloth`.`inbound_cloth_time`" +
		",`xh_inbound_cloth`.`inbound_cloth_note`" +
		",`xh_design`.`design_id`,`xh_design`.`design_name`" +
		",`xh_color`.`color_id`,`xh_color`.`color_name`" +
		",`xh_supplier`.`supplier_id`,`xh_supplier`.`supplier_name`" +
		" FROM `xh_crm`.`xh_inbound`,`xh_crm`.`xh_inbound_cloth`,`xh_crm`.`xh_design`,`xh_crm`.`xh_color`,`xh_crm`.`xh_supplier`" +
		" WHERE `xh_inbound`.`inbound_cloth_id`=`xh_inbound_cloth`.`inbound_cloth_id`" +
		" AND `xh_inbound_cloth`.`design_id`=`xh_design`.`design_id`" +
		" AND `xh_inbound_cloth`.`color_id`=`xh_color`.`color_id`" +
		" AND `xh_inbound_cloth`.`supplier_id`=`xh_supplier`.`supplier_id`" +
		" ORDER BY `inbound_cloth_time` DESC LIMIT ? OFFSET ?"
	log.Println(query2)
	rows, err = p.db.Query(query2, size, offset)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	inboundMap := make(map[string][]inboundInfo)
	for rows.Next() {
		var info inboundInfo
		err = rows.Scan(&info.Id, &info.InboundClothId, &info.Quantity, &info.RemainQuantity, &info.Price, &info.Time, &info.Note,
			&info.DesignId, &info.DesignName, &info.ColorId, &info.ColorName, &info.SupplierId, &info.SupplierName)
		if err != nil {
			log.Println(err)
			continue
		}
		//log.Printf("%+v\n", info)

		inboundInfos, _ := inboundMap[info.Id]
		inboundInfos = append(inboundInfos, info)
		inboundMap[info.Id] = inboundInfos
	}

	var dcMap map[string]Inbound
	lst := make([]InboundInfo, 0)
	for key, value := range inboundMap {
		var inboundInfo InboundInfo
		inboundInfo.Id = key

		dcMap = make(map[string]Inbound)
		for _, v := range value {
			dcId := v.DesignId + v.ColorId
			inbound, ok := dcMap[dcId]
			if ok {
				inbound.TotalQuantity += v.Quantity
				inbound.RemainQuantity += v.RemainQuantity
				dcMap[dcId] = inbound
			} else {
				var info Inbound
				info.Id = v.Id
				info.DesignId = v.DesignId
				info.DesignName = v.DesignName
				info.ColorId = v.ColorId
				info.ColorName = v.ColorName
				info.Price = v.Price
				info.Time = v.Time
				info.Note = v.Note
				dcMap[dcId] = info
			}
		}

		for _, v := range dcMap {
			inboundInfo.Inbounds = append(inboundInfo.Inbounds, v)
		}
		lst = append(lst, inboundInfo)
	}

	return lst, totalCount, nil
}
