package network

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/miekg/dns"
)

// DNSResult ...
type DNSResult struct {
	Successful     bool                 `json:"success"`
	Error          string               `json:"error"`
	Type           string               `json:"type"`
	ServerResponse []*DNSServerResponse `json:"serverResponses"`
}

// DNSServerResponse ...
type DNSServerResponse struct {
	Server string   `json:"server"`
	Msg    *dns.Msg `json:"msg"`
}

// HandleDNSLookup ...
func HandleDNSLookup(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res := DNSResult{
		Successful: true,
		Type:       vars["type"],
	}
	m := new(dns.Msg)

	name := dns.Fqdn(vars["name"])

	switch vars["type"] {

	case "A":
		m.SetQuestion(name, dns.TypeA)
	case "AAAA":
		m.SetQuestion(name, dns.TypeAAAA)
	case "SRV":
		m.SetQuestion(name, dns.TypeSRV)
	case "CNAME":
		m.SetQuestion(name, dns.TypeCNAME)
	case "TXT":
		m.SetQuestion(name, dns.TypeTXT)
	case "MX":
		m.SetQuestion(name, dns.TypeMX)
	case "NS":
		m.SetQuestion(name, dns.TypeNS)
	case "PTR":
		m.SetQuestion(name, dns.TypePTR)
	case "SOA":
		m.SetQuestion(name, dns.TypeSOA)

	}

	// TODO: Get DNS Server from /etc/resolve.conf

	dnsServers := []string{"8.8.8.8:53", "192.168.1.1:53"}

	for _, server := range dnsServers {
		c := new(dns.Client)
		in, _, err := c.Exchange(m, server)

		if err != nil {
			res.Error = err.Error()
			res.Successful = false

		} else {
			srvresp := &DNSServerResponse{
				Server: server,
				Msg:    in,
			}
			res.ServerResponse = append(res.ServerResponse, srvresp)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Failed to write DNS lookup response: %s", err)
	}

}
