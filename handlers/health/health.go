package health

import (
	"fmt"
	"net/http"
)

// Handler to process health check request
func Handler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, "OK")
}
