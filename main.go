package main

import (
	"context"
	"fmt"
	"github.com/sethvargo/go-envconfig"
	"log"
	"net/http"
	"os"
	"time"
)

type Options struct {
	// Port the port to lisetn to
	Port string `env:"PORT,default=8080"`

	// Fail should we fail at all?
	Fail bool `env:"FAIL"`

	// CrashDuration should we periodically fail
	CrashDuration time.Duration `env:"CRASH_DURATION"`
}

func main() {
	o := &Options{}
	ctx := context.TODO()
	err := envconfig.Process(ctx, o)
	if err != nil {
		fmt.Printf("failed: %v\n", err)
		os.Exit(1)
		return
	}

	if o.Fail {
		if o.CrashDuration.Milliseconds() > 0 {
			f := func() {
				fmt.Println("simulating crash now...")
				os.Exit(1)
			}
			fmt.Printf("will crash in %v\n", o.CrashDuration)
			t := time.AfterFunc(o.CrashDuration, f)
			defer t.Stop()
		}
	}
	http.HandleFunc("/", o.handler)

	fmt.Printf("listening to port %s\n", o.Port)
	http.ListenAndServe(":"+o.Port, nil)
}

func (o *Options) handler(w http.ResponseWriter, r *http.Request) {
	title := "Jenkins X golang http example"

	from := ""
	if r.URL != nil {
		from = r.URL.String()
	}
	if from != "/favicon.ico" {
		log.Printf("title: %s\n", title)
	}

	fmt.Fprintf(w, "Hello from:  "+title+"\n")
}
