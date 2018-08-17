package crm

import (
	"errors"
	"log"
	"os"
	"sync"
	"time"
)

type designColorPicCache struct {
	lock      sync.RWMutex
	cachePics map[string]int
}

func newDesignColorPicCache() *designColorPicCache {
	p := new(designColorPicCache)
	p.cachePics = make(map[string]int)
	go p.check()
	return p
}

func (p *designColorPicCache) add(pic string) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.cachePics[pic] = 1
}

func (p *designColorPicCache) remove(pic string) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	_, ok := p.cachePics[pic]
	if ok {
		delete(p.cachePics, pic)
		return nil
	}

	return errors.New("already delete")
}

func (p *designColorPicCache) check() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			p.lock.Lock()
			for pic, ttl := range p.cachePics {
				ttl++
				p.cachePics[pic] = ttl
				if ttl > 15 {
					os.Remove(pic)
					delete(p.cachePics, pic)
				}
			}
			log.Println("cachePics count", len(p.cachePics))
			p.lock.Unlock()
		}
	}
}
