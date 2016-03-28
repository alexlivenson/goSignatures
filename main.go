package main

import (
	"github.com/alexlivenson/signatureCollection/signatures"
	"net/http"
)

var session = signatures.NewSession("signatures")
/*
Create a new MongoDB Session, using a database name "signatures".
Create a new server using that session, then begin listening for
HTTP Requests
*/
func main() {
	context := &signatures.AppContext{session}

	//http.HandleFunc("/signatures", signatures.SignatureHandler(session))
	//mux := mux.NewRouter().StrictSlash(true)

	//http.ListenAndServe(":8090", nil)
	http.ListenAndServe(":8090", signatures.NewSignatureRouter(context))
}
