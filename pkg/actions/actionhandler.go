package actions

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// Data ...
var Data [][]byte

// Message ...
type Message struct {
	Test string
}

// ActionHandler ...
func ActionHandler(rw http.ResponseWriter, r *http.Request) {
	writeResponse := func(response string) {
		if _, err := rw.Write([]byte(response)); err != nil {
			log.Printf("Failed to write action response: %s", err)
		}
	}

	vars := mux.Vars(r)

	switch vars["action"] {

	case "malloc20mb":
		log.Printf("Allocating 20mb to existing %d Mb", len(Data)/2048*2)

		for i := 0; i < 1024*20; i++ {
			kb := make([]byte, 1024)
			if _, err := rand.Read(kb); err != nil {
				http.Error(rw, "Failed to generate memory load data", http.StatusInternalServerError)
				return
			}
			Data = append(Data, kb)
		}

		res := fmt.Sprintf("Size now: %d Mb", len(Data)/2048*2)

		writeResponse(res)

	case "livenessoff":
		//RespondToHealth = false

		writeResponse("Letting /health time out from now on")

	case "fileinfo":
		nofiles := 0
		var size int64
		var files []string
		err := filepath.Walk("/", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			if path == "/dev" || path == "/proc" {
				return filepath.SkipDir
			}
			files = append(files, info.Name())
			nofiles++
			size = size + info.Size()
			return nil
		})
		if err != nil {
			http.Error(rw, "Failed to inspect files", http.StatusInternalServerError)
			return
		}

		res := fmt.Sprintf("Found %d files. Size: %d Mb", nofiles, size/1024/1024)

		writeResponse(res)

	case "log100":
		lines := 100
		start := time.Now()
		for i := 0; i < lines; i++ {
			log.Printf("Logging a lot: %d ", i)

		}
		d := time.Since(start)
		res := fmt.Sprintf("Logged %d lines in %.2f seconds", lines, d.Seconds())

		writeResponse(res)

	case "log1000":
		lines := 1000
		start := time.Now()
		for i := 0; i < lines; i++ {
			log.Printf("Logging a lot: %d ", i)

		}
		d := time.Since(start)
		res := fmt.Sprintf("Logged %d lines in %.2f seconds", lines, d.Seconds())

		writeResponse(res)

	case "log10000":
		lines := 10000
		start := time.Now()
		for i := 0; i < lines; i++ {
			log.Printf("Logging a lot: %d ", i)

		}
		d := time.Since(start)
		res := fmt.Sprintf("Logged %d lines in %.2f seconds", lines, d.Seconds())

		writeResponse(res)

	case "cpusmall":
		const testBytes = `{ "Test": "value" }`
		iter := int64(700000)
		start := time.Now()
		p := &Message{}
		for i := int64(1); i < iter; i++ {
			if err := json.NewDecoder(strings.NewReader(testBytes)).Decode(p); err != nil {
				http.Error(rw, "Failed to process CPU load payload", http.StatusInternalServerError)
				return
			}
		}
		d := time.Since(start)
		res := fmt.Sprintf("[small]. Took %.2f seconds", d.Seconds())
		writeResponse(res)

	case "cpumedium":
		const testBytes = `{ "Test": "value" }`
		iter := int64(3000000)
		start := time.Now()
		p := &Message{}
		for i := int64(1); i < iter; i++ {
			if err := json.NewDecoder(strings.NewReader(testBytes)).Decode(p); err != nil {
				http.Error(rw, "Failed to process CPU load payload", http.StatusInternalServerError)
				return
			}
		}
		d := time.Since(start)
		res := fmt.Sprintf("Done: %.2f s", d.Seconds())
		writeResponse(res)

	case "cpularge":
		const testBytes = `{ "Test": "value" }`
		iter := int64(8000000)
		start := time.Now()
		p := &Message{}
		for i := int64(1); i < iter; i++ {
			if err := json.NewDecoder(strings.NewReader(testBytes)).Decode(p); err != nil {
				http.Error(rw, "Failed to process CPU load payload", http.StatusInternalServerError)
				return
			}
		}
		d := time.Since(start)
		res := fmt.Sprintf("[large]. Took %.2f seconds", d.Seconds())
		writeResponse(res)

		/*case "metrics-increase":
			opsProcessed.Inc()

			rw.Write([]byte("clicks has been increased"))

		case "metrics-gauge-10":
			gauge.Set(10)
			rw.Write([]byte("ata_request_load set to 10"))

		case "metrics-gauge-50":
			gauge.Set(50)
			rw.Write([]byte("ata_request_load set to 50"))

		case "metrics-gauge-90":
			gauge.Set(90)
			rw.Write([]byte("ata_request_load set to 90"))

		case "tracing-flow1":
			span, ctx := opentracing.StartSpanFromContext(r.Context(), "awesome_business_function")
			defer span.Finish()

			time.Sleep(200 * time.Millisecond)

			if !BusinessFunction(ctx) {

				rw.Write([]byte("☠️☠️☠️ Request failed! 🤬 "))

			} else {
				rw.Write([]byte(" 🥳 Request successful! 👻 "))
			}
		*/
	}

}
