package actions

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

var Data [][]byte

type Message struct {
	Test string
}

func ActionHandler(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	switch vars["action"] {

	case "malloc20mb":
		log.Printf("Allocating 20mb to existing %d Mb", len(Data)/2048*2)

		for i := 0; i < 1024*20; i++ {
			kb := make([]byte, 1024)
			rand.Read(kb)
			Data = append(Data, kb)
		}

		res := fmt.Sprintf("Size now: %d Mb", len(Data)/2048*2)

		rw.Write([]byte(res))

	case "livenessoff":
		//RespondToHealth = false

		rw.Write([]byte("Letting /health time out from now on"))

	case "fileinfo":
		nofiles := 0
		var size int64
		var files []string
		filepath.Walk("/", func(path string, info os.FileInfo, err error) error {

			if strings.HasPrefix("/dev", path) {
				return nil
			}
			if strings.HasPrefix("/proc", path) {
				return nil
			}

			if err != nil {
				return nil
			}
			files = append(files, info.Name())
			nofiles++
			size = size + info.Size()
			return nil
		})

		res := fmt.Sprintf("Found %d files. Size: %d Mb", nofiles, size/1024/1024)

		rw.Write([]byte(res))

	case "log100":
		lines := 100
		start := time.Now()
		for i := 0; i < lines; i++ {
			log.Printf("Logging a lot: %d ", i)

		}
		d := time.Since(start)
		res := fmt.Sprintf("Logged %d lines in %.2f seconds", lines, d.Seconds())

		rw.Write([]byte(res))

	case "log1000":
		lines := 1000
		start := time.Now()
		for i := 0; i < lines; i++ {
			log.Printf("Logging a lot: %d ", i)

		}
		d := time.Since(start)
		res := fmt.Sprintf("Logged %d lines in %.2f seconds", lines, d.Seconds())

		rw.Write([]byte(res))

	case "log10000":
		lines := 10000
		start := time.Now()
		for i := 0; i < lines; i++ {
			log.Printf("Logging a lot: %d ", i)

		}
		d := time.Since(start)
		res := fmt.Sprintf("Logged %d lines in %.2f seconds", lines, d.Seconds())

		rw.Write([]byte(res))

	case "cpusmall":
		const testBytes = `{ "Test": "value" }`
		iter := int64(700000)
		start := time.Now()
		p := &Message{}
		for i := int64(1); i < iter; i++ {
			json.NewDecoder(strings.NewReader(testBytes)).Decode(p)
		}
		d := time.Since(start)
		res := fmt.Sprintf("[small]. Took %.2f seconds", d.Seconds())
		rw.Write([]byte(res))

	case "cpumedium":
		const testBytes = `{ "Test": "value" }`
		iter := int64(3000000)
		start := time.Now()
		p := &Message{}
		for i := int64(1); i < iter; i++ {
			json.NewDecoder(strings.NewReader(testBytes)).Decode(p)
		}
		d := time.Since(start)
		res := fmt.Sprintf("Done: %.2f s", d.Seconds())
		rw.Write([]byte(res))

	case "cpularge":
		const testBytes = `{ "Test": "value" }`
		iter := int64(8000000)
		start := time.Now()
		p := &Message{}
		for i := int64(1); i < iter; i++ {
			json.NewDecoder(strings.NewReader(testBytes)).Decode(p)
		}
		d := time.Since(start)
		res := fmt.Sprintf("[large]. Took %.2f seconds", d.Seconds())
		rw.Write([]byte(res))

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
