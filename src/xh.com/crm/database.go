package crm

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

type database struct {
	databaseColor
	databaseCustomer
	databaseSupplier
	databaseDesign
	databaseInboundCloth
	databaseOutboundCloth
	databaseInbound
	databaseOutbound
	databaseScheduling
	databaseDesignColor
	db  *sql.DB
	url string
}

func newDatabase(url string) *database {
	p := new(database)
	p.url = url
	return p
}

// 初始化数据库
func (p *database) connect() error {
	var err error
	p.db, err = sql.Open("mysql", p.url)
	if err == nil {
		p.databaseColor.db = p.db
		p.databaseCustomer.db = p.db
		p.databaseSupplier.db = p.db
		p.databaseDesign.db = p.db
		p.databaseInboundCloth.db = p.db
		p.databaseOutboundCloth.db = p.db
		p.databaseInbound.db = p.db
		p.databaseOutbound.db = p.db
		p.databaseScheduling.db = p.db
		p.databaseDesignColor.db = p.db
	}
	return err
}

func (p *database) updateTaskFlag(flag int, id uint64) error {
	if p.db == nil {
		return errors.New("")
	}
	query := "UPDATE `qk_cloud_media_editor`.`tb_task` SET `flag` = ? WHERE `id` = ?"
	log.Println(query)
	_, err := p.db.Exec(query, flag, id)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (p *database) updateTaskStatusAndFlag(status []byte, flag int, id uint64) error {
	if p.db == nil {
		return errors.New("")
	}
	query := "UPDATE `qk_cloud_media_editor`.`tb_task` SET `status` = ?, `flag` = ? WHERE `id` = ?"
	log.Println(query)
	_, err := p.db.Exec(query, status, flag, id)
	if err != nil {
		log.Println(err)
	}

	return err
}

func (p *database) insertTaskInfo(tasks []string) error {
	if p.db == nil {
		return errors.New("")
	}
	query := "INSERT INTO `qk_cloud_media_editor`.`tb_task` (`task_id`,`task_type`,`task_req`,`status`) VALUES "
	query += strings.Join(tasks, ",") + " ON DUPLICATE KEY UPDATE `update_time`=`update_time`"
	log.Printf("%s\n", query)
	_, err := p.db.Exec(query)
	if err != nil {
		log.Println(err)
	}
	return err
}
