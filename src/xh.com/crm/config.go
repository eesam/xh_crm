package crm

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type config struct {
	Database string
}

func loadConfig(config *config) error {
	filename := "/root/xh_crm/controller.json"
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		return err
	}
	err = json.Unmarshal(b, config)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
