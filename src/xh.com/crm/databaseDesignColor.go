package crm

import (
	"database/sql"
	"errors"
	"log"
)

type DesignColorInfo struct {
	DesignId   string `json:"designId"`
	DesignName string `json:"designName"`
	ColorId    string `json:"colorId"`
	ColorName  string `json:"colorName"`
	PicUrl     string `json:"picUrl"`
	Note       string `json:"note"`
}

type DesignColorList struct {
	TotalCount       int               `json:"totalCount"`
	DesignColorInfos []DesignColorInfo `json:"designColorInfos"`
}

type databaseDesignColor struct {
	db *sql.DB
}

func (p *databaseDesignColor) query(designId string, colorId string) (int, error) {
	if p.db == nil {
		return 0, errors.New("")
	}

	var rows *sql.Rows
	var err error

	query1 := "SELECT count(*) FROM `xh_crm`.`xh_design_color` WHERE `xh_design_color`.`design_id`=? AND `xh_design_color`.`color_id`=?"
	log.Println(query1)
	rows, err = p.db.Query(query1, designId, colorId)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	totalCount := 0
	if rows.Next() {
		err = rows.Scan(&totalCount)
		if err != nil {
			log.Println(err)
			return 0, err
		}
		log.Println("totalCount", totalCount)
	}
	return totalCount, nil
}

func (p *databaseDesignColor) getList(offset int, size int, designId string) ([]DesignColorInfo, int, error) {
	if p.db == nil {
		return nil, 0, errors.New("")
	}

	var rows *sql.Rows
	var err error

	////////////////////////////////////////////////////////
	// totalCount
	query1 := "SELECT count(*) FROM `xh_crm`.`xh_design_color` WHERE `xh_design_color`.`design_id`=?"
	log.Println(query1)
	rows, err = p.db.Query(query1, designId)
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

	query2 := "SELECT `xh_design_color`.`pic_url`,`xh_design_color`.`design_color_note`" +
		",`xh_design_color`.`design_id`,`xh_design`.`design_name`" +
		",`xh_design_color`.`color_id`,`xh_color`.`color_name`" +
		" FROM `xh_crm`.`xh_design_color`,`xh_crm`.`xh_design`,`xh_crm`.`xh_color`" +
		" WHERE `xh_design_color`.`design_id`=?" +
		" AND `xh_design_color`.`design_id`=`xh_design`.`design_id`" +
		" AND `xh_design_color`.`color_id`=`xh_color`.`color_id`" +
		" ORDER BY `design_id` DESC LIMIT ? OFFSET ?"
	log.Println(query2)
	rows, err = p.db.Query(query2, designId, size, offset)
	if err != nil {
		log.Println(err)
		return nil, 0, err
	}

	lst := make([]DesignColorInfo, 0)
	for rows.Next() {
		var info DesignColorInfo
		err = rows.Scan(&info.PicUrl, &info.Note,
			&info.DesignId, &info.DesignName, &info.ColorId, &info.ColorName)
		if err != nil {
			log.Println(err)
			continue
		}

		lst = append(lst, info)
	}
	return lst, totalCount, nil
}

func (p *databaseDesignColor) update(info DesignColorInfo) error {
	if p.db == nil {
		return errors.New("")
	}
	query := "UPDATE `xh_crm`.`xh_design_color` SET `pic_url` = ?, `design_color_note` = ?" +
		" WHERE `design_id` = ? AND `color_id` = ?"
	log.Println(query)
	_, err := p.db.Exec(query, info.PicUrl, info.Note, info.DesignId, info.ColorId)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (p *databaseDesignColor) insert(info DesignColorInfo) error {
	if p.db == nil {
		return errors.New("")
	}
	count, err := p.query(info.DesignId, info.ColorId)
	if err != nil {
		log.Println(err)
		return err
	}

	if count != 0 {
		return p.update(info)
	}
	query := "INSERT INTO `xh_crm`.`xh_design_color` (`design_id`,`color_id`,`pic_url`,`design_color_note`)" +
		" VALUES (?, ?, ?, ?)"
	log.Println(query)
	_, err = p.db.Exec(query, info.DesignId, info.ColorId, info.PicUrl, info.Note)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (p *databaseDesignColor) delete(designId string, colorId string) error {
	if p.db == nil {
		return errors.New("")
	}

	query := "DELETE FROM `xh_crm`.`xh_design_color` WHERE `design_id` = ? AND `color_id` = ?"
	log.Println(query)
	_, err := p.db.Exec(query, designId, colorId)
	if err != nil {
		log.Println(err)
	}

	return nil
}
