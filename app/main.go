package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

type IPResponse struct {
	IP string `json:"ip"`
}

func ipHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-type", "application/json",
	)

	// TODO: use headers if behind proxy!, for instance req.Header.Get("X-FORWARDED-FOR")

	ip, _, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		fmt.Printf("userip: %q is not IP:port", req.RemoteAddr)
		return
	}

	ipRes := IPResponse{
		IP: ip,
	}

	b, err := json.Marshal(ipRes)
	if err != nil {
		fmt.Println("error:", err)
	}

	io.WriteString(res, string(b))
}

func main() {
	listenPort := flag.Int("port", 80, "The port to bind http server")
	listenAddr := flag.String("addr", "", "The addr to bind http server")

	// Parse command line arguments
	flag.Parse()

	http.HandleFunc("/ip", ipHandler)

	bindAddr := fmt.Sprintf("%s:%d", *listenAddr, *listenPort)
	fmt.Println(`Start listenning at "` + bindAddr + `"`)

	log.Fatal(http.ListenAndServe(bindAddr, nil))
}
