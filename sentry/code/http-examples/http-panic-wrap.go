package main

import (
	"fmt"
	"log"
	"net/http"
)

// START OMIT
var tags = map[string]string{`service-name`: `foo`}

type Tester interface {
	DoSomething() error
}

func init() { raven.SetDSN(`<DSN>`) } // HLraven

func main() {
	http.HandleFunc(`/hello-world`, raven.RecoveryHandler(handler)) // HLraven
	log.Fatal(http.ListenAndServe(`:7000`, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	var notImplemented Tester

	if err := notImplemented.DoSomething(); err != nil {
		raven.CaptureError(err, tags) // HLraven
	}

	raven.CaptureMessage(`it's all good`, tags) // HLraven
	fmt.Fprintf(w, `hello world`)
}

// END OMIT
