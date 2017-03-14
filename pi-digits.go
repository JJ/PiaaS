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
	 "encoding/json"
)


func main() {
	http.HandleFunc("/", digits)
	bind := fmt.Sprintf(":%s", os.Getenv("PORT"))
	fmt.Printf("Escucha en %s...", bind)
	err := http.ListenAndServe(bind, nil)
	if err != nil {
		panic(err)
	}
}

func digits(res http.ResponseWriter, req *http.Request) {
	type PiDigits struct {
		Digits     int64
		Pi         string
	}
	digits, _ := strconv.ParseInt(req.URL.Query().Get("digits"), 10, 64)
	this_pi := PiDigits{
		Digits: digits,
		Pi: pigo.Pi(digits),
	}
	js, err := json.Marshal( this_pi )
	if ( err != nil ) {
		fmt.Fprintf( res, "Error %s", err )
	} else {
		res.Header().Set("Content-Type", "application/json")
		res.Write(js)
	}
}
