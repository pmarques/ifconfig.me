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

type IPError struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
}

func ipHandler(res http.ResponseWriter, req *http.Request) {
	log.Println(req.Proto, req.URL)

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
		http.Error(res, "Internal server Error", 500)
		return
	}

	io.WriteString(res, string(b))
}

func main() {
	listenPort := flag.Int("port", 80, "The port to bind http server")
	listenAddr := flag.String("addr", "", "The addr to bind http server")

	// Parse command line arguments
	flag.Parse()

	// Log all the other requests and return 404
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		log.Println(req.Proto, req.URL)

		var errorCode = 404
		var e = IPError{
			Error:   errorCode,
			Message: fmt.Sprintf("Resource [%s] not found", req.URL.Path),
		}
		b, err := json.Marshal(e)
		if err != nil {
			fmt.Println("error:", err)
			http.Error(res, "Internal server Error", 500)
			return
		}

		http.Error(res, string(b), errorCode)
	})
	http.HandleFunc("/ip", ipHandler)

	bindAddr := fmt.Sprintf("%s:%d", *listenAddr, *listenPort)
	fmt.Println(`Start listenning at "` + bindAddr + `"`)

	log.Fatal(http.ListenAndServe(bindAddr, nil))
}
