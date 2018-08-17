package crm

import (
	"database/sql"
	"errors"
	"log"
)

type ColorInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Note string `json:"note"`
}

type ColorList struct {
	TotalCount int         `json:"totalCount"`
	ColorInfos []ColorInfo `json:"colorInfos"`
}

type databaseColor struct {
	db *sql.DB
}

func (p *databaseColor) getList(offset int, size int) ([]ColorInfo, int, error) {
	if p.db == nil {
		return nil, 0, errors.New("")
	}

	var rows *sql.Rows
	var err error

	////////////////////////////////////////////////////////
	// totalCount
	query1 := "SELECT count(*) FROM `xh_crm`.`xh_color`"
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

	query2 := "SELECT `color_id`,`color_name`,`color_note` FROM `xh_crm`.`xh_color` ORDER BY `color_id` DESC LIMIT ? OFFSET ?"
	log.Println(query2)
	rows, err = p.db.Query(query2, size, offset)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	lst := make([]ColorInfo, 0)
	for rows.Next() {
		var info ColorInfo
		err = rows.Scan(&info.Id, &info.Name, &info.Note)
		if err != nil {
			log.Println(err)
			continue
		}

		lst = append(lst, info)
	}
	return lst, totalCount, nil
}

func (p *databaseColor) update(info ColorInfo) error {
	if p.db == nil {
		return errors.New("")
	}
	query := "UPDATE `xh_crm`.`xh_color` SET `color_name` = ?, `color_note` = ? WHERE `color_id` = ?"
	log.Println(query)
	_, err := p.db.Exec(query, info.Name, info.Note, info.Id)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (p *databaseColor) insert(info ColorInfo) error {
	if p.db == nil {
		return errors.New("")
	}
	query := "INSERT INTO `xh_crm`.`xh_color` (`color_id`,`color_name`,`color_note`) VALUES (?, ?, ?)"
	log.Println(query)
	_, err := p.db.Exec(query, info.Id, info.Name, info.Note)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (p *databaseColor) delete(id string) error {
	if p.db == nil {
		return errors.New("")
	}
	query := "DELETE FROM `xh_crm`.`xh_color` WHERE `color_id` = ?"
	log.Println(query)
	_, err := p.db.Exec(query, id)
	if err != nil {
		log.Println(err)
	}
	return err
}
