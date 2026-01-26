package ip

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

// Response represents the response of IP information response
type Response struct {
	XMLName xml.Name `json:"-" xml:"response"`
	IP      string   `json:"ip" xml:"ip"`
}

// Handler to process ip information request
func Handler(res http.ResponseWriter, req *http.Request) {
	log.Println(req.Proto, req.URL)

	var ip string
	var err error

	xffIPs := req.Header.Get("X-FORWARDED-FOR")
	if xffIPs != "" {
		ip = strings.Split(xffIPs, ",")[0]
	} else {
		ip, _, err = net.SplitHostPort(req.RemoteAddr)
		if err != nil {
			http.Error(res, "Error parsing remote address ["+req.RemoteAddr+"]", http.StatusInternalServerError)
			return
		}
	}

	ipRes := Response{
		IP: ip,
	}

	queryString := req.URL.Query()

	var encoding string
	if format, ok := queryString["f"]; ok {
		encoding = format[0]
	} else {
		encoding = "json"
	}

	switch encoding {
	case "json":
		res.Header().Set(
			"Content-type", "application/json",
		)

		b, err := json.Marshal(ipRes)
		if err != nil {
			fmt.Println("error:", err)
			http.Error(res, "Error encoding json", http.StatusInternalServerError)
			return
		}

		res.Write(b)
	case "xml":
		res.Header().Set(
			"Content-type", "application/xml",
		)

		io.WriteString(res, xml.Header)
		enc := xml.NewEncoder(res)
		enc.Indent("  ", "    ")
		if err := enc.Encode(ipRes); err != nil {
			fmt.Printf("error: %v\n", err)
			http.Error(res, "Error encoding xml", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(res, "Encoding response to ["+encoding+"] is not implemented", http.StatusNotImplemented)
	}
}
