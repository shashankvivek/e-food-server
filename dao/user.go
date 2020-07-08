package dao

import (
	"database/sql"
	"e-food/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func FetchUserDetails(db *sql.DB, email string) (*models.RegisterUser, error) {
	row := db.QueryRow("SELECT email,password,firstName,lastName,phoneNo from user_details where email = ?", email)
	userInfo := models.RegisterUser{}

	err := row.Scan(&userInfo.Email,
		&userInfo.Password,
		&userInfo.Fname,
		&userInfo.Lname,
		&userInfo.PhoneNo)
	if err != nil {
		return nil, err
	}
	return &userInfo, nil

}

func RegisterNewUser(db *sql.DB, userInfo *models.RegisterUser) error {
	password := []byte(userInfo.Password)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	row, err := db.Exec("INSERT into customer (email, firstName, lastName, phoneNo, password) values (?,?,?,?,?)", userInfo.Email, userInfo.Fname, userInfo.Lname, userInfo.PhoneNo, hashedPassword)
	if err != nil {
		return err
	}
	if count, _ := row.RowsAffected(); count != 1 {
		return errors.New("Error inserting row value")
	}
	return nil
}
