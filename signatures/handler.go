package signatures

import (
	"encoding/json"
	"net/http"
)

func SignatureIndex(a *AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db := a.Session.DB("signatures")
		result, err := json.Marshal(fetchAllSignatures(db))
		if err != nil {
			panic(err)
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
			panic(err)
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
