package notary

import (
	"fmt"

	"github.com/imdario/mergo"
	log "go.uber.org/zap"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql" // mysql adapter
)

// TODO:config
const (
	UserCollection        = `user`
	UserContacts          = `user_to_user`
	CarContractCollection = `car_contract`
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
//TODO: can't bemaking connections with every request
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

	d.logger.Info(fmt.Sprint("Initializing database ..."))
	// Collection does not exists, let's create it.
	// Execute CREATE TABLE.
	// TODO: This doesn't work with the new script... will have to figure this out
	// if _, err = d.sess.Exec(script); err != nil {
	// 	return nil, err
	// }

	return &d, nil
}

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

	d.logger.Info(fmt.Sprintf("%s added successfully!", collection))
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
// UpdateUser updates user and returns error if any
func (d *dbServer) UpdateUser(user *User) (err error) {
	prev := d.GetUserByUID((*user).UID)
	if prev != nil {
		mergo.MergeWithOverwrite(prev, user)
		res := d.collection(UserCollection).Find("uid = ?", (*user).UID)
		d.logger.Info(fmt.Sprintf("Query created: %v", res))
		defer res.Close()
		err = res.Update(user)
		if err != nil {
			d.logger.Error(fmt.Sprintf("Not cool! %v", err))
		}
	}
	return
}

// LoginUserValidation
func (d *dbServer) LoginUserValidation(user *LoginUser) (*User, error) {

	u := User{}
	res := d.collection(UserCollection).Find("email_address = ?", user.EmailAddress)
	d.logger.Debug(fmt.Sprintf("Query created: %v", res))
	defer res.Close()
	err := res.One(&u)
	if err != nil {
		d.logger.Error(fmt.Sprintf("Not cool! %v", err))
		return nil, err
	}
	d.logger.Info(fmt.Sprintf("User found: %v", u.UID))
	// check if account is verified, and then check if password is valid
	if !u.AccountVerified {
		d.logger.Info(fmt.Sprintf("User '%s' is not verified!", user.EmailAddress))
		return nil, nil
	}
	if !CheckPasswordHash(user.Password, u.Password) {
		d.logger.Info(fmt.Sprintf("User '%s' gave bad password", user.EmailAddress))
		return nil, nil
	}

	return &u, nil
}

//TODO: error
func (d *dbServer) GetUserByUID(uid string) *User {
	u := User{}
	res := d.collection(UserCollection).Find("uid = ?", uid)
	defer res.Close()
	err := res.One(&u)
	if err != nil {
		d.logger.Error(fmt.Sprintf("Not cool! %v", err.Error()))
	}
	return &u
}

//TODO: orderby/ limit?
// SearchUserByName returns users by search string
func (d *dbServer) SearchUserByName(name string) ([]User, error) {
	var users []User
	d.logger.Info(fmt.Sprintf("Search for user by name: %s", name))
	c := fmt.Sprintf("%%%s%%", name)
	cond :=
		db.Or(
			db.Cond{"first_name LIKE": c},
			db.Cond{"last_name LIKE": c})
	res := d.collection(UserCollection).Find(cond)
	defer res.Close()
	err := res.All(&users)
	return users, err
}

func (d *dbServer) FetchUserContacts(uid string) ([]User, error) {
	var users []User
	res := d.sess.Select("*").From(UserCollection).
		Where("uid in (select s_uid from user_to_user where p_uid = ?) ", uid)
	// defer res.Close() TODO: can't figure this out
	err := res.All(&users)
	return users, err
}

func (d *dbServer) GetContractByCID(cid int) *CarContract {
	c := CarContract{}
	res := d.collection(CarContractCollection).Find("cid = ?", cid)
	defer res.Close()
	err := res.One(&c)
	if err != nil {
		d.logger.Error(fmt.Sprintf("Not cool! %v", err.Error()))
	}
	return &c
}

//TODO: This could change everything... edit so it would only change certain fields
// UpdateContract updates the contract and returns error if any
func (d *dbServer) UpdateContract(c *CarContract) (err error) {
	prev := d.GetContractByCID((*c).CID)
	if prev != nil {
		mergo.MergeWithOverwrite(prev, c)
		res := d.collection(CarContractCollection).Find("cid = ?", (*c).CID)
		d.logger.Info(fmt.Sprintf("Query created: %v", res))
		defer res.Close()
		err = res.Update(c)
		if err != nil {
			d.logger.Error(fmt.Sprintf("Not cool! %v", err))
		}
	}
	return
}

func (d *dbServer) FetchUserContracts(uid string) ([]CarContract, error) {
	var c []CarContract
	res := d.sess.Select("*").From(CarContractCollection).
		Where("cid in (select cid from car_contract where buyer=? OR seller=?) ", uid, uid)
	// defer res.Close() TODO: can't figure this out
	err := res.All(&c)
	return c, err
}
