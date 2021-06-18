package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/pmarques/ifconfig.me/handlers/ip"
)

type Error struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
}

func main() {
	listenPort := flag.Int("port", 80, "The port to bind http server")
	listenAddr := flag.String("addr", "", "The addr to bind http server")
	logFileName := flag.String("log-filename", os.Getenv("LOG_FILENAME"), "The filename where to write logs")

	// Parse command line arguments
	flag.Parse()

	if *logFileName != "" {
		log.Println(`Using "` + *logFileName + `" file for logs`)
		file, err := os.OpenFile(*logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		log.SetOutput(file)
	}

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
	log.Println(`Start listenning at "` + bindAddr + `"`)
	fmt.Println(`Start listenning at "` + bindAddr + `"`)

	log.Fatal(http.ListenAndServe(bindAddr, nil))
}
