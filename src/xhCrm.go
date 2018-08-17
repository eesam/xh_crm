package main

import (
	"log"
	"xh.com/crm"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)
	crm.Run()
}
