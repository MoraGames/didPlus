package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"strings"
	"time"
)

type User struct {
	Id int `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	Lastname string `json:"lastName,omitempty"`
	Birthday *time.Time `json:"birthday,omitempty"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"` //TODO: Aggiungere validazione password in "validaros.go"
	db *sql.DB
}

func (u *User) SetDb(db *sql.DB) {
	u.db = db
}

func (u *User) IsValid() error {
	// Create validator object
	v := validator.New()

	// Create custom validator
	//TODO: Da abilitare dopo aver sistemato la regex in validators.go
	//_ = v.RegisterValidation("passwd", validators.ValidatePassword)

	// Validate user object
	err := v.Struct(u)
	if err != nil {
		// Check errors
		var listErrors []string
		for _, e := range err.(validator.ValidationErrors) {
			listErrors = append(listErrors, e.Error())
		}

		// Make single error
		strErr := strings.Join(listErrors, "\n")

		return errors.New(strErr)
	}

	return nil
}

func (u *User) Create() (*User, error) {
	query := "INSERT INTO Users(email, password) VALUES (?, ?)"

	// Create a context
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5 *time.Second)
	defer cancelfunc()

	// Prepare the query
	stmt, err := u.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return nil, err
	}
	defer stmt.Close()

	// Insert the user
	res, err := stmt.ExecContext(ctx, u.Email, u.Password)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return nil, err
	}

	// Get last insert id
	lastId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	// Get the new user info
	u, err = GetUserById(int(lastId), u.db)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func CheckUserExist(email string, db *sql.DB) (bool, error) {
	// Check User
	u, err := GetUserByEmail(email, db)
	if err != nil {
		return false, err
	}

	fmt.Println("DEBUG:CHECK:", u)

	// Check if user email isn't empty
	if u.Email != "" {
		return true, nil
	}

	return false, nil
}

func GetUserByEmail(email string, db *sql.DB) (*User, error){
	query := "SELECT * FROM Users WHERE email=?"
	return getUser(db, query, email)
}

func GetUserById(id int, db *sql.DB) (*User, error) {
	query := "SELECT * FROM Users WHERE id=?"
	return getUser(db, query, id)
}

func getUser(db *sql.DB, query string, args interface{}) (*User, error){
	var u User

	// Prepare query
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	// Do query
	rows, err := stmt.Query(args)
	if err != nil {
		return nil, err
	}

	// Create user object
	for rows.Next() {
		var fn, ln sql.NullString
		err = rows.Scan(&u.Id, &fn, &ln, &u.Birthday, &u.Email, &u.Password)
		if err != nil {
			return nil, err
		}
		u.FirstName = fn.String
		u.Lastname = fn.String
	}

	u.db = db
	return &u, nil
}