package crm

import (
	"database/sql"
	"errors"
	"log"
)

type OutboundClothInfo struct {
	Id             string  `json:"id"`
	InboundClothId string  `json:"inboundClothId"`
	DesignId       string  `json:"designId"`
	DesignName     string  `json:"designName"`
	ColorId        string  `json:"colorId"`
	ColorName      string  `json:"colorName"`
	CustomerId     string  `json:"customerId"`
	CustomerName   string  `json:"customerName"`
	Quantity       float64 `json:"quantity"`
	Price          float64 `json:"price"`
	Time           string  `json:"time"`
	Note           string  `json:"note"`
}

type OutboundClothList struct {
	TotalCount         int                 `json:"totalCount"`
	OutboundClothInfos []OutboundClothInfo `json:"outboundClothInfos"`
}

type databaseOutboundCloth struct {
	db *sql.DB
}

func (p *databaseOutboundCloth) getList(offset int, size int) ([]OutboundClothInfo, int, error) {
	if p.db == nil {
		return nil, 0, errors.New("")
	}

	var rows *sql.Rows
	var err error

	////////////////////////////////////////////////////////
	// totalCount
	query1 := "SELECT count(*) FROM `xh_crm`.`xh_outbound`"
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

	query2 := "SELECT `xh_outbound`.`outbound_id`,`xh_outbound`.`outbound_quantity`,`xh_outbound`.`outbound_price`,`xh_outbound`.`outbound_time`,`xh_outbound`.`outbound_note`" +
		",`xh_outbound`.`design_id`,`xh_design`.`design_name`" +
		",`xh_outbound`.`color_id`,`xh_color`.`color_name`" +
		",`xh_outbound`.`customer_id`,`xh_customer`.`customer_name`" +
		" FROM `xh_crm`.`xh_outbound`,`xh_crm`.`xh_design`,`xh_crm`.`xh_color`,`xh_crm`.`xh_customer`" +
		" WHERE `xh_outbound`.`design_id`=`xh_design`.`design_id`" +
		" AND `xh_outbound`.`color_id`=`xh_color`.`color_id`" +
		" AND `xh_outbound`.`customer_id`=`xh_customer`.`customer_id`" +
		" ORDER BY `outbound_id` DESC LIMIT ? OFFSET ?"
	log.Println(query2)
	rows, err = p.db.Query(query2, size, offset)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	lst := make([]OutboundClothInfo, 0)
	for rows.Next() {
		var info OutboundClothInfo
		err = rows.Scan(&info.Id, &info.Quantity, &info.Price, &info.Time, &info.Note,
			&info.DesignId, &info.DesignName, &info.ColorId, &info.ColorName, &info.CustomerId, &info.CustomerName)
		if err != nil {
			log.Println(err)
			continue
		}

		lst = append(lst, info)
	}
	return lst, totalCount, nil
}

func (p *databaseOutboundCloth) update(info OutboundClothInfo) error {
	if p.db == nil {
		return errors.New("")
	}
	log.Printf("%+v\n", info)
	query := "UPDATE `xh_crm`.`xh_outbound` SET `outbound_quantity` = ?, `outbound_price` = ?, `outbound_time` = ?, `outbound_note` = ?, " +
		"`design_id` = ?, `color_id` = ?, `customer_id` = ?" +
		" WHERE `outbound_id` = ?"
	log.Println(query)
	_, err := p.db.Exec(query, info.Quantity, info.Price, info.Time, info.Note, info.DesignId, info.ColorId, info.CustomerId, info.Id)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (p *databaseOutboundCloth) insert(info OutboundClothInfo, outboundId string) error {
	if p.db == nil {
		return errors.New("")
	}
	log.Printf("%+v\n", info)
	query := "INSERT INTO `xh_crm`.`xh_outbound_cloth` (`outbound_cloth_id`,`outbound_cloth_quantity`,`outbound_cloth_price`,`outbound_cloth_time`,`outbound_cloth_note`," +
		"`inbound_cloth_id`,`customer_id`)" +
		" VALUES (?, ?, ?, ?, ?, ?, ?)"
	log.Println(query)
	_, err := p.db.Exec(query, info.Id, info.Quantity, info.Price, info.Time, info.Note, info.InboundClothId, info.CustomerId)
	if err != nil {
		log.Println(err)
	}

	query = "UPDATE `xh_crm`.`xh_inbound_cloth` SET `remain_quantity` = `remain_quantity`- ?" +
		" WHERE `inbound_cloth_id` = ?"
	log.Println(query)
	_, err = p.db.Exec(query, info.Quantity, info.InboundClothId)
	if err != nil {
		log.Println(err)
	}

	query = "INSERT INTO `xh_crm`.`xh_outbound` (`outbound_id`,`outbound_cloth_id`)" +
		" VALUES (?, ?)"
	_, err = p.db.Exec(query, outboundId, info.Id)
	if err != nil {
		log.Println(err)
	}

	return err
}

func (p *databaseOutboundCloth) delete(id string) error {
	if p.db == nil {
		return errors.New("")
	}
	log.Printf("%+v\n", id)
	query := "DELETE FROM `xh_crm`.`xh_outbound` WHERE `outbound_id` = ?"
	log.Println(query)
	_, err := p.db.Exec(query, id)
	if err != nil {
		log.Println(err)
	}
	return err
}
