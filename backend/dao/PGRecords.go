package dao

import (
	"database/sql"
	"fmt"
)

type AuthPGDatastore struct {
	DB *sql.DB
}

func (apgd *AuthPGDatastore) Init() {
	var err error
	apgd.DB, err = sql.Open("postgres", "user=postgres dbname=auth sslmode=disable")
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
}

func (apgd *AuthPGDatastore) Insert(authRecord *AuthRecord) (*AuthRecord, error) {
	fmt.Println("Entered Insert")
	key := authRecord.UserId
	if len(key) == 0 {
		return nil, fmt.Errorf("invalid key")
	}
	fmt.Println("key: ", key)

	query := "INSERT INTO user_record (userid, password, apikey, apikeyvalid) VALUES ($1, $2, $3, $4) RETURNING userid, password, apikey, apikeyvalid"
	err := apgd.DB.QueryRow(query, authRecord.UserId, authRecord.Password, authRecord.ApiKey, authRecord.ApiKeyValid).Scan(&authRecord.UserId, &authRecord.Password, &authRecord.ApiKey, &authRecord.ApiKeyValid)
	if err != nil {
		return nil, fmt.Errorf("error inserting record: %v", err)
	}
	fmt.Println("Exiting Insert")
	return authRecord, nil
}

func (apgd *AuthPGDatastore) Get(authRecord *AuthRecord) (*AuthRecord, error) {
	fmt.Println("Entered Get")
	var ar AuthRecord
	key := authRecord.UserId
	if len(key) == 0 {
		return nil, fmt.Errorf("invalid key")
	}
	fmt.Println("key: ", key)

	query := "SELECT userid, password, apikey, apikeyvalid FROM user_record WHERE userid = $1"
	err := apgd.DB.QueryRow(query, key).Scan(&ar.UserId, &ar.Password, &ar.ApiKey, &ar.ApiKeyValid)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No record found for key:", key)
			return nil, nil
		}
		return nil, fmt.Errorf("error retrieving record: %v", err)
	}

	fmt.Println("Exiting Get")
	return &ar, nil
}

func (apgd *AuthPGDatastore) Remove(authRecord *AuthRecord) (*AuthRecord, error) {
	fmt.Println("Entered Remove")
	key := authRecord.UserId
	if len(key) == 0 {
		return nil, fmt.Errorf("invalid key")
	}
	fmt.Println("key: ", key)

	ar, err := apgd.Get(authRecord)
	if err != nil {
		return nil, err
	}
	if ar == nil {
		return nil, fmt.Errorf("record does not exist")
	}

	query := "DELETE FROM user_record WHERE userid = $1"
	_, err = apgd.DB.Exec(query, key)
	if err != nil {
		return nil, fmt.Errorf("error removing record: %v", err)
	}

	fmt.Println("Exiting Remove")
	return ar, nil
}

func (apgd *AuthPGDatastore) Update(authRecord *AuthRecord) (*AuthRecord, error) {
	fmt.Println("Entered Update")
	key := authRecord.UserId
	if len(key) == 0 {
		return nil, fmt.Errorf("invalid key")
	}
	fmt.Println("key: ", key)

	ar, err := apgd.Get(authRecord)
	if err != nil {
		return nil, err
	}
	if ar == nil {
		return nil, fmt.Errorf("record does not exist")
	}

	query := "UPDATE user_record SET password = $1, apikey = $2, apikeyvalid = $3 WHERE userid = $4 RETURNING userid, password, apikey, apikeyvalid"
	err = apgd.DB.QueryRow(query, authRecord.Password, authRecord.ApiKey, authRecord.ApiKeyValid, key).Scan(&ar.UserId, &ar.Password, &ar.ApiKey, &ar.ApiKeyValid)
	if err != nil {
		return nil, fmt.Errorf("error updating record: %v", err)
	}

	fmt.Println("Exiting Update")
	return ar, nil
}
