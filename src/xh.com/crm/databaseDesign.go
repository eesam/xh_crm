package crm

import (
	"database/sql"
	"errors"
	"log"
)

type DesignInfo struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	PicUrl   string  `json:"picUrl"`
	Quantity float64 `json:"quantity"`
	Note     string  `json:"note"`
}

type DesignList struct {
	TotalCount  int          `json:"totalCount"`
	DesignInfos []DesignInfo `json:"designInfos"`
}

type databaseDesign struct {
	db *sql.DB
}

func (p *databaseDesign) getList(offset int, size int) ([]DesignInfo, int, error) {
	if p.db == nil {
		return nil, 0, errors.New("")
	}

	var rows *sql.Rows
	var err error

	////////////////////////////////////////////////////////
	// totalCount
	query1 := "SELECT count(*) FROM `xh_crm`.`xh_design`"
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

	query2 := "SELECT `xh_design`.`design_id`,`xh_design`.`design_name`,`xh_design_color`.`pic_url`,`xh_design`.`design_quantity`,`xh_design`.`design_note`" +
		" FROM `xh_crm`.`xh_design`" +
		" LEFT JOIN `xh_crm`.`xh_design_color`" +
		" ON `xh_design`.`design_id`=`xh_design_color`.`design_id`" +
		" GROUP BY `xh_design`.`design_id`" +
		" ORDER BY `design_id` DESC LIMIT ? OFFSET ?"
	log.Println(query2)
	rows, err = p.db.Query(query2, size, offset)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	lst := make([]DesignInfo, 0)
	for rows.Next() {
		var info DesignInfo
		var picUrl []byte
		err = rows.Scan(&info.Id, &info.Name, &picUrl, &info.Quantity, &info.Note)
		if err != nil {
			log.Println(err)
			continue
		}
		info.PicUrl = string(picUrl)
		lst = append(lst, info)
	}
	return lst, totalCount, nil
}

func (p *databaseDesign) update(info DesignInfo) error {
	if p.db == nil {
		return errors.New("")
	}
	query := "UPDATE `xh_crm`.`xh_design` SET `design_name` = ?, `design_quantity` = ?, `design_note` = ? WHERE `design_id` = ?"
	log.Println(query)
	_, err := p.db.Exec(query, info.Name, info.Quantity, info.Note, info.Id)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (p *databaseDesign) insert(info DesignInfo) error {
	if p.db == nil {
		return errors.New("")
	}
	query := "INSERT INTO `xh_crm`.`xh_design` (`design_id`,`design_name`,`design_quantity`,`design_note`) VALUES (?, ?, ?, ?)"
	log.Println(query)
	_, err := p.db.Exec(query, info.Id, info.Name, info.Quantity, info.Note)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (p *databaseDesign) delete(id string) error {
	if p.db == nil {
		return errors.New("")
	}
	query := "DELETE FROM `xh_crm`.`xh_design` WHERE `design_id` = ?"
	log.Println(query)
	_, err := p.db.Exec(query, id)
	if err != nil {
		log.Println(err)
	}
	return err
}
