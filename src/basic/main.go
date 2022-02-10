package main

// middleware pattern

import "log"

type Trace interface {
	Log()
}

type traceService struct {
	name string
}

func (t traceService) Log() {
	log.Println("name: ", t.name)
}

type chainPre struct {
	name string
	next Trace
}

func (cp chainPre) Log() {
	log.Println("name: ", cp.name)
	// call next func
	cp.next.Log()
}

func main() {
	var svc Trace
	svc = traceService{"traceService"}
	svc = chainPre{"chainPre", svc}
	svc.Log()
}
