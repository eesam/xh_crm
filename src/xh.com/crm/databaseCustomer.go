package crm

import (
	"database/sql"
	"errors"
	"log"
)

type CustomerInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Note string `json:"note"`
}

type CustomerList struct {
	TotalCount    int            `json:"totalCount"`
	CustomerInfos []CustomerInfo `json:"customerInfos"`
}

type databaseCustomer struct {
	db *sql.DB
}

func (p *databaseCustomer) getList(offset int, size int) ([]CustomerInfo, int, error) {
	if p.db == nil {
		return nil, 0, errors.New("")
	}

	var rows *sql.Rows
	var err error

	////////////////////////////////////////////////////////
	// totalCount
	query1 := "SELECT count(*) FROM `xh_crm`.`xh_customer`"
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

	query2 := "SELECT `customer_id`,`customer_name`,`customer_note` FROM `xh_crm`.`xh_customer` ORDER BY `customer_id` DESC LIMIT ? OFFSET ?"
	log.Println(query2)
	rows, err = p.db.Query(query2, size, offset)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	lst := make([]CustomerInfo, 0)
	for rows.Next() {
		var info CustomerInfo
		err = rows.Scan(&info.Id, &info.Name, &info.Note)
		if err != nil {
			log.Println(err)
			continue
		}

		lst = append(lst, info)
	}
	return lst, totalCount, nil
}

func (p *databaseCustomer) update(info CustomerInfo) error {
	if p.db == nil {
		return errors.New("")
	}
	query := "UPDATE `xh_crm`.`xh_customer` SET `customer_name` = ?, `customer_note` = ? WHERE `customer_id` = ?"
	log.Println(query)
	_, err := p.db.Exec(query, info.Name, info.Note, info.Id)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (p *databaseCustomer) insert(info CustomerInfo) error {
	if p.db == nil {
		return errors.New("")
	}
	query := "INSERT INTO `xh_crm`.`xh_customer` (`customer_id`,`customer_name`,`customer_note`) VALUES (?, ?, ?)"
	log.Println(query)
	_, err := p.db.Exec(query, info.Id, info.Name, info.Note)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (p *databaseCustomer) delete(id string) error {
	if p.db == nil {
		return errors.New("")
	}
	query := "DELETE FROM `xh_crm`.`xh_customer` WHERE `customer_id` = ?"
	log.Println(query)
	_, err := p.db.Exec(query, id)
	if err != nil {
		log.Println(err)
	}
	return err
}
