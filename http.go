package main

import (
	"os"
    "fmt"
    "net/http"
    "log"
	"io/ioutil"
	"strings"
)

func linux(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("/etc/system-release")
	if err != nil {
		//Do something
	}
	s := []string{string(content), "on", os.Getenv("HOSTNAME")}
    fmt.Fprintf(w, strings.Join(s, " ")) // send data to client side
}

func main() {
    http.HandleFunc("/", linux) // set router
    err := http.ListenAndServe(":80", nil) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}