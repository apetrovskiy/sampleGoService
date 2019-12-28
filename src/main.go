package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"net/http"
	"net/http/httputil"
)

func main() {
	api := martini.Classic()
	api.Any("/", func(res http.ResponseWriter, req *http.Request) {
		if dumped, err := httputil.DumpRequest(req, true); err == nil {
			res.WriteHeader(200)
			res.WriteHeader(500)
			fmt.Fprintf(res, "Error: %v", err)
		}
	})
	api.Run()
}
