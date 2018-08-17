package crm

import (
	"database/sql"
	"errors"
	"log"
)

type InboundClothInfo struct {
	Id             string  `json:"id"`
	DesignId       string  `json:"designId"`
	DesignName     string  `json:"designName"`
	ColorId        string  `json:"colorId"`
	ColorName      string  `json:"colorName"`
	SupplierId     string  `json:"supplierId"`
	SupplierName   string  `json:"supplierName"`
	Quantity       float64 `json:"quantity"`
	RemainQuantity float64 `json:"remainQuantity"`
	Price          float64 `json:"price"`
	Time           string  `json:"time"`
	Note           string  `json:"note"`
}

type InboundClothList struct {
	TotalCount        int                `json:"totalCount"`
	InboundClothInfos []InboundClothInfo `json:"inboundClothInfos"`
}

type databaseInboundCloth struct {
	db *sql.DB
}

func (p *databaseInboundCloth) getList(designId string, colorId string) ([]InboundClothInfo, int, error) {
	if p.db == nil {
		return nil, 0, errors.New("")
	}

	var rows *sql.Rows
	var err error

	////////////////////////////////////////////////////////
	// totalCount
	query1 := "SELECT count(*) FROM `xh_crm`.`xh_inbound_cloth` WHERE `xh_inbound_cloth`.`design_id`=? AND `xh_inbound_cloth`.`color_id`=?"
	log.Println(query1)
	rows, err = p.db.Query(query1, designId, colorId)
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

	query2 := "SELECT `xh_inbound_cloth`.`inbound_cloth_id`,`xh_inbound_cloth`.`inbound_cloth_quantity`,`xh_inbound_cloth`.`remain_quantity`" +
		",`xh_inbound_cloth`.`inbound_cloth_price`" +
		",`xh_inbound_cloth`.`inbound_cloth_time`,`xh_inbound_cloth`.`inbound_cloth_note`" +
		",`xh_design`.`design_id`,`xh_design`.`design_name`" +
		",`xh_color`.`color_id`,`xh_color`.`color_name`" +
		",`xh_supplier`.`supplier_id`,`xh_supplier`.`supplier_name`" +
		" FROM `xh_crm`.`xh_inbound_cloth`,`xh_crm`.`xh_design`,`xh_crm`.`xh_color`,`xh_crm`.`xh_supplier`" +
		" WHERE `xh_inbound_cloth`.`design_id`=? AND `xh_inbound_cloth`.`color_id`=?" +
		" AND `xh_inbound_cloth`.`design_id`=`xh_design`.`design_id`" +
		" AND `xh_inbound_cloth`.`color_id`=`xh_color`.`color_id`" +
		" AND `xh_inbound_cloth`.`supplier_id`=`xh_supplier`.`supplier_id`" +
		" ORDER BY `inbound_cloth_time` DESC"
	log.Println(query2)
	rows, err = p.db.Query(query2, designId, colorId)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	lst := make([]InboundClothInfo, 0)
	for rows.Next() {
		var info InboundClothInfo
		err = rows.Scan(&info.Id, &info.Quantity, &info.RemainQuantity, &info.Price, &info.Time, &info.Note,
			&info.DesignId, &info.DesignName, &info.ColorId, &info.ColorName, &info.SupplierId, &info.SupplierName)
		if err != nil {
			log.Println(err)
			continue
		}

		lst = append(lst, info)
	}
	return lst, totalCount, nil
}

func (p *databaseInboundCloth) update(info InboundClothInfo) error {
	if p.db == nil {
		return errors.New("")
	}
	log.Printf("%+v\n", info)
	query := "UPDATE `xh_crm`.`xh_inbound_cloth` SET `inbound_quantity` = ?, `inbound_price` = ?, `inbound_time` = ?, `inbound_note` = ?, " +
		"`design_id` = ?, `color_id` = ?, `supplier_id` = ?" +
		" WHERE `inbound_id` = ?"
	log.Println(query)
	_, err := p.db.Exec(query, info.Quantity, info.Price, info.Time, info.Note, info.DesignId, info.ColorId, info.SupplierId, info.Id)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (p *databaseInboundCloth) insert(info InboundClothInfo, inboundId string) error {
	if p.db == nil {
		return errors.New("")
	}
	log.Printf("%+v\n", info)
	query := "INSERT INTO `xh_crm`.`xh_inbound_cloth` (`inbound_cloth_id`,`inbound_cloth_quantity`,`remain_quantity`,`inbound_cloth_price`,`inbound_cloth_time`,`inbound_cloth_note`," +
		"`design_id`,`color_id`,`supplier_id`)" +
		" VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	log.Println(query)
	_, err := p.db.Exec(query, info.Id, info.Quantity, info.RemainQuantity, info.Price, info.Time, info.Note, info.DesignId, info.ColorId, info.SupplierId)
	if err != nil {
		log.Println(err)
	}

	query = "INSERT INTO `xh_crm`.`xh_inbound` (`inbound_id`,`inbound_cloth_id`)" +
		" VALUES (?, ?)"
	_, err = p.db.Exec(query, inboundId, info.Id)
	if err != nil {
		log.Println(err)
	}

	return err
}

func (p *databaseInboundCloth) delete(id string) error {
	if p.db == nil {
		return errors.New("")
	}
	log.Printf("%+v\n", id)
	query := "DELETE FROM `xh_crm`.`xh_inbound_cloth` WHERE `inbound_cloth_id` = ?"
	log.Println(query)
	_, err := p.db.Exec(query, id)
	if err != nil {
		log.Println(err)
	}
	return err
}
