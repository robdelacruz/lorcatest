package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/zserge/lorca"
)

func main() {
	err := run(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
func run(args []string) error {
	lorcaargs := []string{}
	ui, err := lorca.New("", "", 480, 320, lorcaargs...)
	if err != nil {
		return err
	}
	defer ui.Close()

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	defer ln.Close()
	fmt.Printf("Listening on %s...\n", ln.Addr())
	go http.Serve(ln, http.FileServer(http.Dir("./www")))
	ui.Load(fmt.Sprintf("http://%s/", ln.Addr()))

	/*
		http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./www"))))
		port := "8000"
		fmt.Printf("Listening on %s...\n", port)
		go http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
		ui.Load("http://localhost:8000")
	*/

	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}

	fmt.Printf("Exiting.\n")
	return nil
}
