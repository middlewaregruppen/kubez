package api

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Delay struct {
	DelayMin int `json:"delayMin"`
	DelayMax int `json:"delayMax"`
}

func (d *Delay) DelayRequest() {
	// Calulate how long time to delay.
	delay := 0
	if d.DelayMin < d.DelayMax {
		delay = rand.Intn(d.DelayMax-d.DelayMin) + d.DelayMin
	}

	// Delay
	if delay > 0 {
		d, err := time.ParseDuration(fmt.Sprintf("%dms", delay))
		if err != nil {
			log.Print(err)
		}
		time.Sleep(d)
	}

}
