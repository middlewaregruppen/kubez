package main

import (
	"crypto/rand"
	"log"
	"time"
)

var Data [][]byte

func main() {

	for i := 0; i < 50; i++ {
		log.Printf("Allocating 10mb to existing %d Mb", len(Data)/2048*2)
		for i := 0; i < 1024*10; i++ {
			kb := make([]byte, 1024)
			rand.Read(kb)
			Data = append(Data, kb)
		}
		log.Printf("Size now: %d Mb", len(Data)/2048*2)
		time.Sleep(time.Second)
	}

	log.Printf("Done! Sleeping forever")

	for {
		time.Sleep(time.Second)
	}

}
