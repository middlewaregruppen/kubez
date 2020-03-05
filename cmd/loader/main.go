package main

import (
	"crypto/rand"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	petname "github.com/dustinkirkland/golang-petname"
)

func main() {

	var profile string
	var megaBytes int
	var logwait time.Duration
	flag.CommandLine.StringVar(&profile, "profile", "none", "The loadprofile to run")
	flag.CommandLine.IntVar(&megaBytes, "mb", 100, "MB to allocate if load profile is mem")
	flag.CommandLine.DurationVar(&logwait, "logwait", time.Millisecond*10, "duration to wait between writing log lines.")

	flag.Parse()

	switch profile {

	case "cpu":
		cpu()
	case "mem":
		mem(megaBytes)
	case "log":
		logger(logwait)
	case "none":
		none()

	}

}

// logger logs heaps of data
func logger(logwait time.Duration) {
	type LogItem struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Data string `json:"data"`
	}

	for {
		start := time.Now()
		for i := 0; i < 1000; i++ {
			name := petname.Generate(1, "")
			data := petname.Generate(4, " haz ")
			li := LogItem{
				ID:   i,
				Name: name,
				Data: data,
			}
			ll, _ := json.Marshal(li)
			log.Printf("%s", ll)

			time.Sleep(logwait)
		}

		duration := time.Since(start)

		li := LogItem{
			ID:   1000,
			Name: "STATUS",
			Data: fmt.Sprintf("Logged 1000 lines in %f seconds", duration.Seconds()),
		}
		ll, _ := json.Marshal(li)
		log.Printf("%s", ll)

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
