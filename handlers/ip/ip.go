package ip

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

// Response represents the response of IP information response
type Response struct {
	XMLName xml.Name `json:"-" xml:"response"`
	IP      string   `json:"ip" xml:"ip"`
}

// Handler to process ip information request
func Handler(res http.ResponseWriter, req *http.Request) {
	log.Println(req.Proto, req.URL)

	// TODO: use headers if behind proxy!, for instance req.Header.Get("X-FORWARDED-FOR")

	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		fmt.Printf("userip: %q is not IP:port", req.RemoteAddr)
		return
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
			http.Error(res, "Internal server Error", 500)
			return
		}

		io.WriteString(res, string(b))
	case "xml":
		res.Header().Set(
			"Content-type", "application/xml",
		)

		io.WriteString(res, xml.Header)
		enc := xml.NewEncoder(res)
		enc.Indent("  ", "    ")
		if err := enc.Encode(ipRes); err != nil {
			fmt.Printf("error: %v\n", err)
			http.Error(res, "Internal server Error", 500)
		}
	default:
		http.Error(res, "Encoding responso to ["+encoding+"] is not implemented", http.StatusNotImplemented)
	}
}
