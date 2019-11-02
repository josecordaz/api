package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	Name  string `json:"name"`
	EMail string `json:"email"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		p := person{
			Name:  "Jose Carlos",
			EMail: "cordaz@pinpoint.com",
		}

		bts, _ := json.Marshal(p)

		fmt.Fprint(w, string(bts))
	})

	fmt.Println("Server running")
	if err := http.ListenAndServe(":8283", nil); err != nil {
		log.Fatal(err)
	}
}
