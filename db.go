package notary

import (
	"fmt"

	log "go.uber.org/zap"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql" // mysql adapter
)

// TODO:config
const (
	UserCollection   = `user`
	UserContacts     = `user_to_user`
	CarContract      = `car_contract`
	CarContractUsers = `car_contract_user`
)

type dbServer struct {
	sess   sqlbuilder.Database
	logger *log.Logger
}

func (d *dbServer) CloseSession() {
	d.sess.Close()
}

func (d *dbServer) GetSession() sqlbuilder.Database {
	return d.sess
}

//TODO: might want to add a Ping() just to check whether db conn is alive
//TODO: configure max open connections/idle connections, etc...

func dbInit(script string, settings mysql.ConnectionURL, logger *log.Logger) (*dbServer, error) {
	var d dbServer
	var err error
	d.logger = logger

	// var s sqlbuilder.Database
	d.sess, err = mysql.Open(settings)
	if err != nil {
		return nil, err
	}

	// Collection lookup.
	col := d.sess.Collection(`ntry`)
	if col.Exists() {
		return &d, nil
	}

	logger.Info(fmt.Sprint("Initializing database ...", UserCollection))
	// Collection does not exists, let's create it.
	// Execute CREATE TABLE.
	// TODO: This doesn't work with the new script... will have to figure this out
	// if _, err = d.sess.Exec(script); err != nil {
	// 	return nil, err
	// }

	return &d, nil
}

/**
 * DB wrapper using context
 */
// type dbwrapper struct {
// 	h         http.Handler
// 	dbSession *dbServer
// }

// func WithDB(s *dbServer, h http.Handler) http.Handler {
// 	return &dbwrapper{dbSession: s, h: h}
// }

// func (dbw *dbwrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	log.Println("In db wrapper, Server HTTP")
// 	// // copy the Session
// 	// dbcopy := dbw.dbSession
// 	// defer dbcopy.Close()

// 	// context.Set(r, "db", dbcopy)
// 	dbw.h.ServeHTTP(w, r)
// }

// Returns true if collection already exist
func (d *dbServer) CollectionExist(name string) bool {
	return d.sess.Collection(name).Exists()
}

func (d *dbServer) CreateCollection(name, script string) error {
	if _, err := d.sess.Exec(script); err != nil {
		d.logger.Error(fmt.Sprintf("Couldn't create table! should be Shutting down..."))
		return err
	}
	d.logger.Info(fmt.Sprint("Collection created successfully: ", name))
	return nil
}

//TODO: might want to add salt to the pwd
// InsertUser inserts new row in the database, in the given collection
func (d *dbServer) Insert(anything interface{}, collection string) error {
	_, err := d.collection(collection).Insert(anything)
	if err != nil {
		d.logger.Error(fmt.Sprintf("Oops! User couldn't be added! %v", err))
		return err
	}

	d.logger.Info(fmt.Sprint("User added successfully: ", anything))
	return nil
}

// UserExistsByUniqueField returns boolean whether user exists by particular field value
func (d *dbServer) UserExistsByUniqueField(user *User) (bool, error) {

	res := d.collection(UserCollection).Find("uid = ? OR email_address = ?", (*user).UID, (*user).EmailAddress)
	defer res.Close()
	if count, err := res.Count(); err != nil {
		d.logger.Error(fmt.Sprintf("Not cool! %v", err))
		return false, err
	} else if count > 0 {
		return true, nil
	}
	return false, nil
}

func (d *dbServer) collection(collection string) db.Collection {
	return d.sess.Collection(collection)
}

//TODO: This could change everything... edit so it would only change certain fields
// UpdateUser returns boolean whether user was updated or not
func (d *dbServer) UpdateUser(user *User) (err error) {

	res := d.collection(UserCollection).Find("uid = ?", user.UID)
	defer res.Close()
	err = res.Update(user)
	if err != nil {
		d.logger.Error(fmt.Sprintf("Not cool! %v", err))
	}
	return
}

// LoginUserValidation
func (d *dbServer) LoginUserValidation(user *LoginUser) (*User, error) {

	u := User{}
	res := d.collection(UserCollection).Find("password = ? AND email_address = ?", (*user).Password, (*user).EmailAddress)
	d.logger.Info(fmt.Sprintf("Query created: %v", res))
	defer res.Close()
	err := res.One(&u)
	if err != nil {
		d.logger.Error(fmt.Sprintf("Not cool! %v", err))
		return nil, err
	}
	d.logger.Info(fmt.Sprintf("User: %v", u))
	if !u.AccountVerified {
		return nil, nil
	}
	return &u, nil
}

func (d *dbServer) GetUserByUID(uid string) *User {
	u := User{}
	res := d.collection(UserCollection).Find("uid = ?", uid)
	defer res.Close()
	err := res.Update(&u)
	if err != nil {
		d.logger.Error(fmt.Sprintf("Not cool! %v", err.Error()))
	}
	return &u
}

// // GetUserValidationCode
// func (d *dbServer) GetUserValidationCode(user *VerifyUserSignature) string {
// 	u := User{}
// 	res := d.collection(defaultUserCollection).Find("secondary_address = ?", (*user).PubKey)
// 	defer res.Close()
// 	err := res.Select("verification_code").One(&u)
// 	if err != nil {
// 		log.Println("Not cool!", err)
// 	}
// 	return u.VerificationCode
// }
