package signatures

import "labix.org/v2/mgo"

/*
Want to use different database for tests, so embed *mgo.Session
and store the database name
 */
type DatabaseSession struct {
	*mgo.Session
	databaseName string
}

/*
Connect to the local MongoDB and set up the database.
*/
func NewSession(name string) *DatabaseSession {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	addIndexToSignatureEmails(session.DB(name))
	return &DatabaseSession{session, name}
}

/*
Add a unique index on the "email" field. This doesn't prevent users from signing twice,
since they can still enter "dudebro+signature2@exmaple.com".
 */
func addIndexToSignatureEmails(db *mgo.Database) {
	index := mgo.Index{
		Key: 	  []string{"email"},
		Unique:   true,
		DropDups: true,
	}
	err := db.C("signatures").EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}