package signatures_test

import (
	"testing"
	"net/http"
	"github.com/alexlivenson/signatureCollection/signatures"
)

var dbName = "test_signatures"

func TestGetSignatures(t *testing.T) {
	// Given
	session := signatures.NewSession(dbName)
	db := session.DB("dbName")
	err := db.C(dbName).Insert(signatures.Signature{
		FirstName: "Jim",
		LastName: "Bob",
	})

	// When
	http.HandleFunc("/signatures", signatures.SignatureHandler(session))
	req, err := http.NewRequest("GET", "/signatures", nil)

	if err != nil {
		t.Error(err)
	}
}
