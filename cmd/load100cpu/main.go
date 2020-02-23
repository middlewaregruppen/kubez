package main

import (
	"encoding/json"
	"log"
	"strings"
	"time"
)

type Message struct {
	Test string
}

func main() {
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
		time.Sleep(5 * time.Millisecond)
	}
}
