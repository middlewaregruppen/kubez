package network

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/miekg/dns"
)

type DNSResult struct {
	Successful     bool                 `json:"success"`
	Error          string               `json:"error"`
	Type           string               `json:"type"`
	ServerResponse []*DNSServerResponse `json:"serverResponses"`
}

type DNSServerResponse struct {
	Server string   `json:"server"`
	Msg    *dns.Msg `json:"msg"`
}

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
		break
	case "AAAA":
		m.SetQuestion(name, dns.TypeAAAA)
		break
	case "SRV":
		m.SetQuestion(name, dns.TypeSRV)
		break
	case "CNAME":
		m.SetQuestion(name, dns.TypeCNAME)
		break
	case "TXT":
		m.SetQuestion(name, dns.TypeTXT)
		break
	case "MX":
		m.SetQuestion(name, dns.TypeMX)
		break
	case "NS":
		m.SetQuestion(name, dns.TypeNS)
		break
	case "PTR":
		m.SetQuestion(name, dns.TypePTR)
		break
	case "SOA":
		m.SetQuestion(name, dns.TypeSOA)
		break

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

	b, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)

}
