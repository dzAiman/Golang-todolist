// https://levelup.gitconnected.com/build-a-todo-app-in-golang-mongodb-and-react-e1357b4690a6

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"./route"
)

func main() {
	fmt.Printf("Hello world")
	// log.Info("starting go server")
	rr := route.Router()
	// router.HandleFunc("/helloworld", HW).Methods("GET")
	http.ListenAndServe(":8000", rr)
	log.Fatal(http.ListenAndServe(":8080", rr))

}

func HW(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"here I am!": true}`)
}

func init() {
	// log.SetFormatter(&log.TextFormatter{})
	// log.SetReportCaller(true)
}
