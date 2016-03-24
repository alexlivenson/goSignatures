package main

import (
	"github.com/alexlivenson/signatureCollection/signatures"
	"net/http"
)

/*
Create a new MongoDB Session, using a database name "signatures".
Create a new server using that session, then begin listening for
HTTP Requests
 */
func main() {
	session := signatures.NewSession("signatures")
	http.HandleFunc("/signatures", signatures.SignatureHandler(session))
	http.ListenAndServe(":8090", nil)
}
