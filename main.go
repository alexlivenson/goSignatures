package main

import (
	"github.com/alexlivenson/signatureCollection/signatures"
	"github.com/gorilla/mux"
	"net/http"
	"text/template"
)

var session = signatures.NewSession("signatures")
var templates = template.Must(template.ParseFiles("public/views/index.html"))

/*
Create a new MongoDB Session, using a database name "signatures".
Create a new server using that session, then begin listening for
HTTP Requests
*/
func main() {
	context := &signatures.AppContext{session}
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").
		Path("/").
		Name("IndexHandler").
		HandlerFunc(signatures.Logger(IndexHandler, "IndexHandler"))
	signatures.AppendSignatureRouter(router, context)
	http.ListenAndServe(":8090", router)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
