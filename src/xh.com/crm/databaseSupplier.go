package crm

import (
	"database/sql"
	"errors"
	"log"
)

type SupplierInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Note string `json:"note"`
}

type SupplierList struct {
	TotalCount    int            `json:"totalCount"`
	SupplierInfos []SupplierInfo `json:"supplierInfos"`
}

type databaseSupplier struct {
	db *sql.DB
}

func (p *databaseSupplier) getList(offset int, size int) ([]SupplierInfo, int, error) {
	if p.db == nil {
		return nil, 0, errors.New("")
	}

	var rows *sql.Rows
	var err error

	////////////////////////////////////////////////////////
	// totalCount
	query1 := "SELECT count(*) FROM `xh_crm`.`xh_supplier`"
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

	query2 := "SELECT `supplier_id`,`supplier_name`,`supplier_note` FROM `xh_crm`.`xh_supplier` ORDER BY `supplier_id` DESC LIMIT ? OFFSET ?"
	log.Println(query2)
	rows, err = p.db.Query(query2, size, offset)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	lst := make([]SupplierInfo, 0)
	for rows.Next() {
		var info SupplierInfo
		err = rows.Scan(&info.Id, &info.Name, &info.Note)
		if err != nil {
			log.Println(err)
			continue
		}

		lst = append(lst, info)
	}
	return lst, totalCount, nil
}

func (p *databaseSupplier) update(info SupplierInfo) error {
	if p.db == nil {
		return errors.New("")
	}
	query := "UPDATE `xh_crm`.`xh_supplier` SET `supplier_name` = ?, `supplier_note` = ? WHERE `supplier_id` = ?"
	log.Println(query)
	_, err := p.db.Exec(query, info.Name, info.Note, info.Id)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (p *databaseSupplier) insert(info SupplierInfo) error {
	if p.db == nil {
		return errors.New("")
	}
	query := "INSERT INTO `xh_crm`.`xh_supplier` (`supplier_id`,`supplier_name`,`supplier_note`) VALUES (?, ?, ?)"
	log.Println(query)
	_, err := p.db.Exec(query, info.Id, info.Name, info.Note)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (p *databaseSupplier) delete(id string) error {
	if p.db == nil {
		return errors.New("")
	}
	query := "DELETE FROM `xh_crm`.`xh_supplier` WHERE `supplier_id` = ?"
	log.Println(query)
	_, err := p.db.Exec(query, id)
	if err != nil {
		log.Println(err)
	}
	return err
}
