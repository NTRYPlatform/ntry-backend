package main

import (
	"io/ioutil"
	"log"

	db "upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

var (
	c              = GetDatabaseSettings()
	ntrydb         *sqlbuilder.Database
	userCollection *db.Collection
	//TODO: configure max open connections/idle connections, etc...
	settings = mysql.ConnectionURL{
		Host:     c.Host,
		Database: c.Name,
		User:     c.User,
		Password: c.Password,
	}
)

//TODO: will have to figure out something about the closing
func initConnection() {
	if ntrydb == nil {
		sess, err := mysql.Open(settings)
		if err != nil {
			log.Println("Database connection cannot be made!")
			log.Fatalf("db.Open(): %q\n", err)
		} else {
			log.Println("Connected to db using the following DSN: ", settings.String())
		}
		//check for existence of reqd collections
		if sess.Collection("user").Exists() == false {
			log.Println("User table doesn't exist! Let's see if we can create it...")

			bytes, err := ioutil.ReadFile("config/ntry-user.sql")
			if err != nil {
				log.Panicln("Can't read sql file!", err)
			}
			_, err = sess.Exec(string(bytes))
			if err != nil {
				log.Fatalf("Couldn't create table! Shutting down...")
			}

		}

		ntrydb = &sess
	} //TODO: might want to add a Ping() just to check whether db conn is alive
}

func initUserCollection() {
	if userCollection == nil {
		coll := (*ntrydb).Collection("user")
		userCollection = &coll
	}
}

//TODO: might want to add salt to the pwd
// InsertUser inserts new user in the database
func InsertUser(user User) bool {
	initConnection()
	initUserCollection()
	added := false
	// coll := (*ntrydb).Collection("user")
	// _, err := coll.Insert(user)
	_, err := (*userCollection).Insert(user)
	if err != nil {
		log.Println("Oops! User couldn't be added!", err)
	} else {
		log.Println("User %v added successfully!", user)
		added = true
	}
	return added
}

// UserExistsByUniqueField returns boolean whether user exists by particular field value
func UserExistsByUniqueField(user *User) bool {
	initConnection()
	initUserCollection()
	exists := false
	res := (*userCollection).Find("eth_address = ? OR email_address = ?", (*user).EthAddress, (*user).EmailAddress)
	count, err := res.Count()
	if err != nil {
		log.Println("Not cool!", err)
	}
	if count > 0 {
		exists = true
	}
	return exists
}

// UserExistsBy
func LoginUserValidation(user *LoginUser) *User {
	initConnection()
	initUserCollection()
	u := User{}
	res := (*userCollection).Find("password = ? AND email_address = ?", (*user).Password, (*user).EmailAddress)
	err := res.One(&u)
	if err != nil {
		log.Println("Not cool!", err)
	}
	return &u
}
