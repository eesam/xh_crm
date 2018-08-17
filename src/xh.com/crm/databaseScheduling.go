package crm

import (
	"database/sql"
	"errors"
	"log"
)

type SchedulingInfo struct {
	Id           string  `json:"id"`
	DesignId     string  `json:"designId"`
	DesignName   string  `json:"designName"`
	ColorId      string  `json:"colorId"`
	ColorName    string  `json:"colorName"`
	SupplierId   string  `json:"supplierId"`
	SupplierName string  `json:"supplierName"`
	Quantity     float64 `json:"quantity"`
	Price        float64 `json:"price"`
	Time         string  `json:"time"`
	Note         string  `json:"note"`
}

type SchedulingList struct {
	TotalCount      int              `json:"totalCount"`
	SchedulingInfos []SchedulingInfo `json:"schedulingInfos"`
}

type databaseScheduling struct {
	db *sql.DB
}

func (p *databaseScheduling) getList(offset int, size int) ([]SchedulingInfo, int, error) {
	if p.db == nil {
		return nil, 0, errors.New("")
	}

	var rows *sql.Rows
	var err error

	////////////////////////////////////////////////////////
	// totalCount
	query1 := "SELECT count(*) FROM `xh_crm`.`xh_scheduling`"
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

	query2 := "SELECT `xh_scheduling`.`scheduling_id`,`xh_scheduling`.`scheduling_quantity`,`xh_scheduling`.`scheduling_price`,`xh_scheduling`.`scheduling_time`,`xh_scheduling`.`scheduling_note`" +
		",`xh_scheduling`.`design_id`,`xh_design`.`design_name`" +
		",`xh_scheduling`.`color_id`,`xh_color`.`color_name`" +
		",`xh_scheduling`.`supplier_id`,`xh_supplier`.`supplier_name`" +
		" FROM `xh_crm`.`xh_scheduling`,`xh_crm`.`xh_design`,`xh_crm`.`xh_color`,`xh_crm`.`xh_supplier`" +
		" WHERE `xh_scheduling`.`design_id`=`xh_design`.`design_id`" +
		" AND `xh_scheduling`.`color_id`=`xh_color`.`color_id`" +
		" AND `xh_scheduling`.`supplier_id`=`xh_supplier`.`supplier_id`" +
		" ORDER BY `scheduling_id` DESC LIMIT ? OFFSET ?"
	log.Println(query2)
	rows, err = p.db.Query(query2, size, offset)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	lst := make([]SchedulingInfo, 0)
	for rows.Next() {
		var info SchedulingInfo
		err = rows.Scan(&info.Id, &info.Quantity, &info.Price, &info.Time, &info.Note,
			&info.DesignId, &info.DesignName, &info.ColorId, &info.ColorName, &info.SupplierId, &info.SupplierName)
		if err != nil {
			log.Println(err)
			continue
		}

		lst = append(lst, info)
	}
	return lst, totalCount, nil
}

func (p *databaseScheduling) update(info SchedulingInfo) error {
	if p.db == nil {
		return errors.New("")
	}
	query := "UPDATE `xh_crm`.`xh_scheduling` SET `scheduling_quantity` = ?, `scheduling_price` = ?, `scheduling_time` = ?, `scheduling_note` = ?, " +
		"`design_id` = ?, `color_id` = ?, `supplier_id` = ?" +
		" WHERE `scheduling_id` = ?"
	log.Println(query)
	_, err := p.db.Exec(query, info.Quantity, info.Price, info.Time, info.Note, info.DesignId, info.ColorId, info.SupplierId, info.Id)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (p *databaseScheduling) insert(info SchedulingInfo) error {
	if p.db == nil {
		return errors.New("")
	}
	query := "INSERT INTO `xh_crm`.`xh_scheduling` (`scheduling_id`,`scheduling_quantity`,`scheduling_price`,`scheduling_time`,`scheduling_note`," +
		"`design_id`,`color_id`,`supplier_id`)" +
		" VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	log.Println(query)
	_, err := p.db.Exec(query, info.Id, info.Quantity, info.Price, info.Time, info.Note, info.DesignId, info.ColorId, info.SupplierId)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (p *databaseScheduling) delete(id string) error {
	if p.db == nil {
		return errors.New("")
	}
	query := "DELETE FROM `xh_crm`.`xh_scheduling` WHERE `scheduling_id` = ?"
	log.Println(query)
	_, err := p.db.Exec(query, id)
	if err != nil {
		log.Println(err)
	}
	return err
}
