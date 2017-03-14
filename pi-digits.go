package main
// Taken from https://play.golang.org/p/mb5eoZpnYN
// No license specified

import (
	"fmt"
	"net/http"
	"os"
//	"runtime"
	"github.com/JJ/pigo"
	"strconv"
)

func main() {
	http.HandleFunc("/", digits)
	bind := fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"), os.Getenv("OPENSHIFT_GO_PORT"))
	fmt.Printf("Escucha en %s...", bind)
	err := http.ListenAndServe(bind, nil)
	if err != nil {
		panic(err)
	}
}

func digits(res http.ResponseWriter, req *http.Request) {
	digits, _ := strconv.ParseInt(req.URL.Query().Get("digits"), 10, 64)
	fmt.Fprintf(res, "%d dígitos de Pi: %s", digits, pigo.Pi(digits))
}
