package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"strings"
	"time"
)

var statsd = &StatsD{Namespace: "leftpad", SampleRate: 0.5}

func init() {
	var f, err = ioutil.TempFile("", "leftpad.log")
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)
}

func leftpad(s string, length int, char rune) string {
	for len(s) < length {
		s = string(char) + s
	}
	return s
}

type leftpadResponse struct {
	Str string `json:"str"`
}

func timedHandler(name string, nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		nextFunc(w, r)
		elapsed := time.Since(start)
		statsd.Timing(fmt.Sprintf("request.%s.timing", name), elapsed)
		statsd.Incr(fmt.Sprintf("request.%s.count", name))
		log.Printf("%s request took %v", name, elapsed)
	}
}

func leftpadHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("url %s\tip %s\tua %s", r.RequestURI, r.RemoteAddr, r.UserAgent())
	str := r.FormValue("str")
	length, err := strconv.Atoi(r.FormValue("len"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	statsd.Histogram("request.str.len", float64(length))
	chr := ' '
	if len(r.FormValue("chr")) > 0 {
		chr = []rune(r.FormValue("chr"))[0]
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp := leftpadResponse{Str: leftpad(str, length, chr)}
	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type StatsD struct {
	Namespace  string
	SampleRate float64
}

func (s *StatsD) Send(stat string, kind string, delta float64) {
	buf := fmt.Sprintf("%s.", s.Namespace)
	trimmedStat := strings.NewReplacer(":", "_", "|", "_", "@", "_").Replace(stat)
	buf += fmt.Sprintf("%s:%s|%s", trimmedStat, delta, kind)
	if s.SampleRate != 0 && s.SampleRate < 1 {
		buf += fmt.Sprintf("|@%s", strconv.FormatFloat(s.SampleRate, 'f', -1, 64))
	}
	ioutil.Discard.Write([]byte(buf)) // TODO: Write to a socket
}

func (s *StatsD) Incr(stat string) {
	s.Send(stat, "c", 1)
}

func (s *StatsD) Histogram(stat string, value float64) {
	s.Send(stat, "h", value)
}

func (s *StatsD) Timing(stat string, value time.Duration) {
	s.Send(stat, "ms", value.Seconds()*1000)
}

func main() {
	http.HandleFunc("/v1/leftpad/", timedHandler("leftpad", leftpadHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
