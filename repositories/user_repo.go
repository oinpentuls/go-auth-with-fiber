package repositories

import (
	"database/sql"
	"errors"

	"github.com/oinpentuls/auth-with-fiber/database"
	"github.com/oinpentuls/auth-with-fiber/models"
	"golang.org/x/crypto/bcrypt"
)

func StoreUser(user models.User) (models.User, error) {
	db := database.DB

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	res, err := db.Exec(
		`INSERT INTO users (name, email, password) 
		VALUES (?, ?, ?)`,
		user.Name, user.Email, password)

	if err != nil {
		return user, err
	}

	lastId, _ := res.LastInsertId()

	row := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", lastId)

	err = row.Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		return user, err
	}
	return user, err
}

func GetUserByEmail(email string) (models.User, error) {
	db := database.DB
	var user models.User

	row := db.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?", email)

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err == sql.ErrNoRows {
		return user, errors.New("email not found with any users")
	}

	if err != nil {
		return user, err
	}

	return user, nil
}

func GetUserById(id string) (models.User, error) {
	db := database.DB
	var user models.User

	row := db.QueryRow("SELECT id, name, email, password FROM users WHERE id = ?", id)

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err == sql.ErrNoRows {
		return user, errors.New("user not found!")
	}

	if err != nil {
		return user, err
	}

	return user, nil
}
