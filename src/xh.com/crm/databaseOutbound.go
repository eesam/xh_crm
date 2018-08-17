package crm

import (
	"database/sql"
	"errors"
	"log"
)

type outboundInfo struct {
	Id              string
	OutboundClothId string
	DesignId        string
	DesignName      string
	ColorId         string
	ColorName       string
	Quantity        float64
	RemainQuantity  float64
	Price           float64
	Time            string
	Note            string
}

type Outbound struct {
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

type OutboundInfo struct {
	Id        string     `json:"id"`
	Outbounds []Outbound `json:"outbounds"`
}

type OutboundList struct {
	TotalCount    int            `json:"totalCount"`
	OutboundInfos []OutboundInfo `json:"outboundInfos"`
}

type databaseOutbound struct {
	db *sql.DB
}

func (p *databaseOutbound) getList(offset int, size int) ([]OutboundInfo, int, error) {
	if p.db == nil {
		return nil, 0, errors.New("")
	}

	var rows *sql.Rows
	var err error

	////////////////////////////////////////////////////////
	// totalCount
	query1 := "SELECT count(DISTINCT `outbound_id`) FROM `xh_crm`.`xh_outbound`"
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

	query2 := "SELECT `xh_outbound`.`outbound_id`" +
		",`xh_outbound_cloth`.`outbound_cloth_id`" +
		",`xh_outbound_cloth`.`outbound_cloth_quantity`" +
		",`xh_inbound_cloth`.`remain_quantity`" +
		",`xh_outbound_cloth`.`outbound_cloth_price`" +
		",`xh_outbound_cloth`.`outbound_cloth_time`" +
		",`xh_outbound_cloth`.`outbound_cloth_note`" +
		",`xh_design`.`design_id`,`xh_design`.`design_name`" +
		",`xh_color`.`color_id`,`xh_color`.`color_name`" +
		" FROM `xh_crm`.`xh_outbound`,`xh_crm`.`xh_outbound_cloth`,`xh_crm`.`xh_design`,`xh_crm`.`xh_color`,`xh_crm`.`xh_inbound_cloth`" +
		" WHERE `xh_outbound`.`outbound_cloth_id`=`xh_outbound_cloth`.`outbound_cloth_id`" +
		" AND `xh_outbound_cloth`.`inbound_cloth_id`=`xh_inbound_cloth`.`inbound_cloth_id`" +
		" AND `xh_inbound_cloth`.`design_id`=`xh_design`.`design_id`" +
		" AND `xh_inbound_cloth`.`color_id`=`xh_color`.`color_id`" +
		" ORDER BY `outbound_cloth_time` DESC LIMIT ? OFFSET ?"
	log.Println(query2)
	rows, err = p.db.Query(query2, size, offset)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	outboundMap := make(map[string][]outboundInfo)
	for rows.Next() {
		var info outboundInfo
		err = rows.Scan(&info.Id, &info.OutboundClothId, &info.Quantity, &info.RemainQuantity, &info.Price, &info.Time, &info.Note,
			&info.DesignId, &info.DesignName, &info.ColorId, &info.ColorName)
		if err != nil {
			log.Println(err)
			continue
		}
		//log.Printf("%+v\n", info)

		outbounds, _ := outboundMap[info.Id]
		outbounds = append(outbounds, info)
		outboundMap[info.Id] = outbounds
	}

	var dcMap map[string]Outbound
	lst := make([]OutboundInfo, 0)
	for key, value := range outboundMap {
		var outboundInfo OutboundInfo
		outboundInfo.Id = key

		dcMap = make(map[string]Outbound)
		for _, v := range value {
			dcId := v.DesignId + v.ColorId
			outbound, ok := dcMap[dcId]
			if ok {
				outbound.TotalQuantity += v.Quantity
				outbound.RemainQuantity += v.RemainQuantity
				dcMap[dcId] = outbound
			} else {
				var info Outbound
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
			outboundInfo.Outbounds = append(outboundInfo.Outbounds, v)
		}
		lst = append(lst, outboundInfo)
	}

	return lst, totalCount, nil
}
