package main

import (
	"context"
	"fmt"
	"github.com/jenkins-x/jx-logging/v3/pkg/log"
	"github.com/sethvargo/go-envconfig"
	"go.uber.org/atomic"
	"net/http"
	"os"
	"time"
)

type Options struct {
	// Port the port to listen to
	Port string `env:"PORT,default=8080"`

	// Crash should we crash the process at the CrashDuration
	Crash bool `env:"CRASH"`

	// CrashDuration should we periodically fail
	CrashDuration time.Duration `env:"CRASH_DURATION"`

	// RequestFailCount how often should we fail a request. Zero means no failures, otherwise its every X requests we return a fail
	RequestFailCount int `env:"REQUEST_FAIL"`

	// RequestErrorCode the HTTP code returned when failing http requests
	RequestErrorCode int `env:"REQUEST_ERROR_CODE,default=404"`

	requestCounter atomic.Int32
}

func main() {
	log.Logger().Infof("starting up")

	o := &Options{}
	ctx := context.TODO()
	err := envconfig.Process(ctx, o)
	if err != nil {
		fmt.Printf("failed: %v\n", err)
		os.Exit(1)
		return
	}

	if o.Crash {
		if o.CrashDuration.Milliseconds() > 0 {
			f := func() {
				log.Logger().Infof("simulating crash now...")
				os.Exit(1)
			}
			log.Logger().Infof("will crash in %v\n", o.CrashDuration)
			t := time.AfterFunc(o.CrashDuration, f)
			defer t.Stop()
		}
	}
	http.HandleFunc("/", o.handler)

	log.Logger().Infof("listening to port %s\n", o.Port)
	http.ListenAndServe(":"+o.Port, nil)
}

func (o *Options) handler(w http.ResponseWriter, r *http.Request) {
	m := o.RequestFailCount
	if m > 0 {
		c := int(o.requestCounter.Add(1))
		if c%m == 0 {
			log.Logger().Errorf("failing request with %v", o.RequestErrorCode)
			w.WriteHeader(o.RequestErrorCode)
			w.Write([]byte("simulating failure of service"))
			return
		}
	}

	title := "Jenkins X golang http example"

	from := ""
	if r.URL != nil {
		from = r.URL.String()
	}
	if from != "/favicon.ico" {
		log.Logger().Infof("title: %s\n", title)
	}

	fmt.Fprintf(w, "Hello from:  "+title+"\n")
}
