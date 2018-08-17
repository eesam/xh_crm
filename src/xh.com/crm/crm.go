package crm

import (
	"log"
)

type crm struct {
	config              config
	db                  *database
	httpSvr             *httpServer
	designColorPicCache *designColorPicCache
}

var instance *crm

func Run() {
	instance = newInstance()
	instance.run()
}

func newInstance() *crm {
	p := new(crm)
	return p
}

func (p *crm) run() {
	// 读配置文件
	err := loadConfig(&p.config)
	if err != nil {
		log.Fatalln(err)
	}

	p.db = newDatabase(p.config.Database)
	err = p.db.connect()
	if err != nil {
		log.Fatalln(err)
	}

	p.designColorPicCache = newDesignColorPicCache()

	p.httpSvr = newHttpServer()
	p.httpSvr.run()
}
