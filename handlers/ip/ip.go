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

const (
	formatJSON = "json"
	formatXML  = "xml"
)

// getClientIP retrieves the client's IP address from the request.
// It checks the "X-FORWARDED-FOR" header first, then falls back to RemoteAddr.
func getClientIP(req *http.Request) (string, error) {
	xffIPs := req.Header.Get("X-FORWARDED-FOR")
	if xffIPs != "" {
		return strings.Split(xffIPs, ",")[0], nil
	}
	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		return "", fmt.Errorf("Error parsing remote address [%s]: %w", req.RemoteAddr, err)
	}
	return ip, nil
}

// Response represents the response of IP information response
type Response struct {
	XMLName xml.Name `json:"-" xml:"response"`
	IP      string   `json:"ip" xml:"ip"`
}

// Handler to process ip information request
func Handler(res http.ResponseWriter, req *http.Request) {
	log.Println(req.Proto, req.URL)

	ip, err := getClientIP(req)
	if err != nil {
		log.Printf("Error getting client IP: %v", err)
		http.Error(res, "Error retrieving client IP", http.StatusInternalServerError)
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
		encoding = formatJSON // Default to JSON
	}

	switch encoding {
	case formatJSON:
		res.Header().Set(
			"Content-type", "application/json",
		)

		b, err := json.Marshal(ipRes)
		if err != nil {
			fmt.Println("error:", err)
			http.Error(res, "Error encoding json", http.StatusInternalServerError)
			return
		}

		io.WriteString(res, string(b))
	case formatXML:
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
		http.Error(res, fmt.Sprintf("Encoding response to [%s] is not implemented", encoding), http.StatusNotImplemented)
	}
}
