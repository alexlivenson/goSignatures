package signatures

import (
	"net/http"
	"fmt"
	"encoding/json"
)

var dbName = "signatures"
/*
Create a new server. We'll use JSON rendered and MongoDB
Database handler. We define two routes: "GET /signatures"
and "POST /signatures"
 */
func NewServer(session *DatabaseSession) {
	// Create the server and set up middleware
	http.HandleFunc("/signatures", func(w http.ResponseWriter, r *http.Request) {
		db := session.DB(dbName)
		switch r.Method {
		case "GET":
			result, err := json.Marshal(fetchAllSignatures(db))
			if err != nil {
				fmt.Println(err)
				return
			}
			w.Write([]byte(string(result)))
		case "POST":
			decoder := json.NewDecoder(r.Body)
			var sig Signature
			err := decoder.Decode(&sig)
			if err != nil {
				panic(err)
			}

			if sig.valid() {
				err := db.C(dbName).Insert(sig)
				if err == nil {
					w.WriteHeader(http.StatusCreated)
				} else {
					w.WriteHeader(http.StatusBadRequest)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})
	http.ListenAndServe(":8080", nil)
}

func SignatureHandler(session *DatabaseSession) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db := session.DB("signatures")
		switch r.Method {
		case "GET":
			result, err := json.Marshal(fetchAllSignatures(db))
			if err != nil {
				fmt.Println(err)
				return
			}
			w.Write([]byte(string(result)))
		case "POST":
			decoder := json.NewDecoder(r.Body)
			var sig Signature
			err := decoder.Decode(&sig)
			if err != nil {
				panic(err)
			}

			if sig.valid() {
				err := db.C("signatures").Insert(sig)
				if err == nil {
					w.WriteHeader(http.StatusCreated)
				} else {
					w.WriteHeader(http.StatusBadRequest)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}