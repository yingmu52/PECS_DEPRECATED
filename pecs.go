package main

import (
	"PECS/solution"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", showResult)
	http.ListenAndServe(":8080", nil)
}

func showResult(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.pecs")
	if err != nil {
		log.Fatalln(err)
	}
	circles := solution.SimmulatedAnnealing(6)
	circlesForDisplay := circles.ConvertToCanvas(400)
	fmt.Println("done")
	t.Execute(w, circlesForDisplay)
}
