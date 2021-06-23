package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

func main() {
	port := flag.Int("p", 8090, "port")
	flag.Parse()

	url, err := url.Parse("http://localhost:8080")
	if err != nil {
		panic(err)
	}

	director1 := func(req *http.Request) {
		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
	}
	director2 := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = "localhost:8081"
	}

	reverseProxy1 := &httputil.ReverseProxy{Director: director1}
	reverseProxy2 := &httputil.ReverseProxy{Director: director2}
	handler := handler{proxy: []*httputil.ReverseProxy{reverseProxy1, reverseProxy2}}

	http.Handle("/", handler)

	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}

type handler struct {
	proxy []*httputil.ReverseProxy
}

// handler used inside the http.Handle() should implement this ServreHTTP() function
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header.Get("x-forwarded-for"))
	port, err := strconv.Atoi(r.Header.Get("x-forwarded-for"))
	if err != nil {
		w.WriteHeader(500)
		return
	}
	port = port % 10
	if len(h.proxy) < port+1 {
		w.WriteHeader(404)
		return
	}
	h.proxy[port].ServeHTTP(w, r)
}
