package signatures

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"labix.org/v2/mgo/bson"
	"net/http"
)

func SignatureIndex(a *AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db := a.Session.DB("signatures")
		result, err := json.Marshal(fetchAllSignatures(db))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write([]byte(string(result)))
	}
}

func SignatureCreate(a *AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var sig Signature

		err := decoder.Decode(&sig)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		if sig.valid() {
			collection := a.Session.DB(a.Session.databaseName).C("signatures")
			err := collection.Insert(sig)
			if err == nil {
				w.WriteHeader(http.StatusCreated)
			} else {
				w.WriteHeader(http.StatusBadRequest)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func SignatureShow(a *AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if id := vars["id"]; bson.IsObjectIdHex(id) {
			db := a.Session.DB("signatures")

			result, err := json.Marshal(findSignatureById(bson.ObjectIdHex(id), db))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Write([]byte(string(result)))
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func SignatureDelete(a *AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if id := vars["id"]; bson.IsObjectIdHex(id) {
			db := a.Session.DB("signatures")

			err := removeSignatureById(bson.ObjectIdHex(id), db)
				if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
