package api

import (
	"io"
	"net/http"

	"github.com/juju/ratelimit"
)

type TransmissionRate struct {
	ResponseRate int64 `json:"responseRate"`
	RequestRate  int64 `json:"requestRate"`
}

func (tr *TransmissionRate) RequestReader(r *http.Request) io.Reader {
	if tr.RequestRate == 0 {
		return r.Body
	}

	bucket := ratelimit.NewBucketWithRate(float64(tr.RequestRate), tr.RequestRate)
	return ratelimit.Reader(r.Body, bucket)

}

func (tr *TransmissionRate) ResponseReader(response io.Reader) (reader io.Reader, buffSize int64) {

	bucket := ratelimit.NewBucketWithRate(float64(tr.ResponseRate), tr.ResponseRate)
	return ratelimit.Reader(response, bucket), tr.ResponseRate

}
