package main

import (
	"crypto/rand"
	"encoding/json"
	"flag"
	"log"
	"strings"
	"time"
)

func main() {

	var profile string
	var megaBytes int
	flag.CommandLine.StringVar(&profile, "profile", "none", "The loadprofile to run")
	flag.CommandLine.IntVar(&megaBytes, "mb", 100, "MB to allocate if load profile is mem")

	flag.Parse()

	switch profile {

	case "cpu":
		cpu()
	case "mem":
		mem(megaBytes)
	case "none":
		none()

	}

}

// cpu loads the cpu with 100% until stopped.
func cpu() {
	type Message struct {
		Test string
	}

	log.Print("Loading CPU 100%")
	for {
		const testBytes = `{ "Test": "value" }`
		iter := int64(8000000)
		start := time.Now()
		p := &Message{}
		for i := int64(1); i < iter; i++ {
			json.NewDecoder(strings.NewReader(testBytes)).Decode(p)
		}
		d := time.Since(start)

		log.Printf("[Load cycle complete]. Took %.2f seconds", d.Seconds())
		time.Sleep(1 * time.Millisecond)
	}

}

// mem allocates memory to megaBytes and then sleeps until killed
func mem(megaBytes int) {
	var Data [][]byte

	log.Printf("Allocating memory to %d MB", megaBytes)

	for megaBytes > (len(Data) / 2048 * 2) {
		log.Printf("Allocating 10mb to existing %d Mb", len(Data)/2048*2)
		for i := 0; i < 1024*10; i++ {
			kb := make([]byte, 1024)
			rand.Read(kb)
			Data = append(Data, kb)
		}
		log.Printf("Size now: %d Mb", len(Data)/2048*2)

		time.Sleep(time.Second * 2)
	}

	log.Printf("Done! Sleeping forever")

	for {
		time.Sleep(time.Second)
	}

}

// none does nothing
func none() {
	log.Print("I do nothing!")
	for {
		time.Sleep(time.Second)
	}
}
