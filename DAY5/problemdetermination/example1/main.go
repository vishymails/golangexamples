package main

import (
	"net/http"
	"net/http/pprof"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Profiling"))
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", helloHandler)

	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdLine", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	http.ListenAndServe(":8080", r)

}
