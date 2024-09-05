package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
    mux := http.NewServeMux()
    addr := flag.String("addr", ":4000", "HTTP network address");
    flag.Parse()
    fmt.Println(addr);
    //flag.Parse();
    /* serves files over http 
    ∙ serve files using the provided directory as root
    */
    fileServer := http.FileServer(http.Dir("./ui/static/"));
    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet/view", snippetView)
    mux.HandleFunc("/snippet/create", snippetCreate);
    mux.Handle("/static/", http.StripPrefix("/static", fileServer));
    
    log.Println("Starting server on :4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
