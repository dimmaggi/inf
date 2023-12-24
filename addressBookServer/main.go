package main

import (
	"Addressprj/controllers/stdhttp"
	"Addressprj/gate/psg"
	"fmt"
	"net/http"
)

func main() {
	s, err := psg.NewPsg("postgres://postres:667765@localhost:5432/addressbook", "postgres", "667765")
	if err != nil {
		fmt.Println("Error")
	}
	stdhttp.NewController("localhost:8080", s)
	http.ListenAndServe(":8080", nil)
}
