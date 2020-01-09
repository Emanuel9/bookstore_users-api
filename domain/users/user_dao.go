package users

import (
	"fmt"
	"github.com/Emanuel9/bookstore_users-api/datasources/mysql/users_db"
	"github.com/Emanuel9/bookstore_users-api/utils/date_utils"
	"github.com/Emanuel9/bookstore_users-api/utils/errors"
	"strings"
)

const(
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows = "no rows in result set"
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, data_created) VALUES(?, ?, ?, ?); "
	queryGetUser = "SELECT id, first_name, last_name, email, data_created FROM users where id =?;"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to get user %d: %s", user.Id, err.Error()))
	}


	//
	//result := usersDB[user.Id]
	//if result == nil {
	//	return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	//}
	//
	//user.Id = result.Id
	//user.FirstName = result.FirstName
	//user.LastName = result.LastName
	//user.Email = result.Email
	//user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s is already taken", user.Email))
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	user.Id = userId
	//
	//current := usersDB[user.Id]
	//if current != nil {
	//	if current.Email == user.Email {
	//		return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
	//	}
	//	return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	//}
	//
	//user.DateCreated = date_utils.GetNowString()
	//usersDB[user.Id] = user
	return nil
}