package users

import (
	"github.com/Emanuel9/bookstore_users-api/datasources/mysql/users_db"
	"github.com/Emanuel9/bookstore_users-api/utils/date_utils"
	"github.com/Emanuel9/bookstore_users-api/utils/errors"
	"github.com/Emanuel9/bookstore_users-api/utils/mysql_utils"
)

const(
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, data_created) VALUES(?, ?, ?, ?); "
	queryGetUser = "SELECT id, first_name, last_name, email, data_created FROM users where id =?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
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
		return mysql_utils.ParseError(err)
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
		return mysql_utils.ParseError(err)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
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

func (user *User) Update() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}