// signatures/signature.go

package signatures

import "labix.org/v2/mgo"

/*
Each signature is composed of first name, last name,
email, age, and a short message. When represented in JSON, ditch
TitleCase for snake_case
 */

type Signature struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int `json:"age"`
	Message   string `json:"message"`
}

/*
I want to make sure all these fields are present. The message is optional
, but if present should be less then 140 characters
 */
func (s *Signature) valid() bool {
	return len(s.FirstName) > 0 &&
	len(s.LastName) > 0 &&
	len(s.Email) > 0 &&
	s.Age >= 18 && s.Age <= 180 &&
	len(s.Message) < 140
}
/*
I'll use this method when displaying all signatures for "GET /signatures"
. Consult the mgo docs for more info: http://godoc.org/labix.org/v2/mgo
 */
func fetchAllSignatures(db *mgo.Database) []Signature {
	signatures := []Signature{}
	err := db.C("signatures").Find(nil).All(&signatures)

	if err != nil {
		panic(err)
	}
	return signatures
}