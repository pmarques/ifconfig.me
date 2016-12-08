package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/pmarques/ifconfig.me/handlers/ip"
)

type Error struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
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
		var e = Error{
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
	http.HandleFunc("/ip", ip.Handler)

	bindAddr := fmt.Sprintf("%s:%d", *listenAddr, *listenPort)
	fmt.Println(`Start listenning at "` + bindAddr + `"`)

	log.Fatal(http.ListenAndServe(bindAddr, nil))
}
