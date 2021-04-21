package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/sethvargo/go-envconfig"
)

type Options struct {
	// Port the port to lisetn to
	Port string `env:"PORT,default=8080"`
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

	http.HandleFunc("/", o.handler)

	fmt.Printf("listening to port %s\n", o.Port)
	http.ListenAndServe(":" + o.Port, nil)
}


func (o*Options) handler(w http.ResponseWriter, r *http.Request) {
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

