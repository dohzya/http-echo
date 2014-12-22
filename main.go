package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	httpHost := flag.String("h", "127.0.0.1", "The host to bind")
	httpPort := flag.Int("p", 8080, "The port to listen")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.Copy(w, r.Body)
	})

	hostPort := fmt.Sprintf("%s:%d", *httpHost, *httpPort)
	fmt.Printf("Listen on %v\n", hostPort)
	log.Fatal(http.ListenAndServe(hostPort, nil))
}
