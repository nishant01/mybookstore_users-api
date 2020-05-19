package users

import (
	"fmt"

	"github.com/nishant01/mybookstore_users-api/logger"

	"github.com/nishant01/mybookstore_users-api/datasources/mysql/users_db"

	"github.com/nishant01/mybookstore_users-api/utils/errors"
)

const (
	INDEX_UNIQUE_EMAIL        = "email_UNIQUE"
	NO_ROW_IN_RESULT          = "no rows in result set"
	QUERY_INSERT_USER         = "INSERT INTO users(first_name, last_name, email, status, password, date_created) VALUES(?,?,?,?,?,?);"
	QUERY_GET_USER_BY_ID      = "SELECT id, first_name, last_name, email, status, date_created FROM users WHERE id = ?;"
	QUERY_UPDATE_USER         = "UPDATE users SET first_name=?, last_name=?, email=?, status=?, password=? WHERE id=?;"
	QUERY_DELETE_USER         = "DELETE FROM users WHERE id=?;"
	QUERY_FIND_USER_BY_STATUS = "SELECT id, first_name, last_name, email, status, date_created FROM users WHERE status=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(QUERY_GET_USER_BY_ID)
	if err != nil {
		logger.Error("error while trying to get user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.DateCreated); getErr != nil {
		logger.Error("error while trying to get user by id", getErr)
		return errors.NewInternalServerError("database error")
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(QUERY_INSERT_USER)
	if err != nil {
		logger.Error("error while trying to save user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.Password, user.DateCreated)

	if saveErr != nil {
		logger.Error("error while trying to save user", saveErr)
		return errors.NewInternalServerError("database error")
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewNotFoundError(fmt.Sprintf("error to save user %s", err.Error()))
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(QUERY_UPDATE_USER)
	if err != nil {
		logger.Error("error while trying to update user", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	_, updateErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.Password, user.Id)
	if updateErr != nil {
		logger.Error("error while trying to update user", updateErr)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(QUERY_DELETE_USER)
	if err != nil {
		logger.Error("error while trying to delete user", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	_, deleteErr := stmt.Exec(user.Id)
	if deleteErr != nil {
		logger.Error("error while trying to delete user", deleteErr)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(QUERY_FIND_USER_BY_STATUS)
	if err != nil {
		logger.Error("error while trying to find by status", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	rows, findErr := stmt.Query(status)
	if findErr != nil {
		return nil, errors.NewInternalServerError(findErr.Error())
	}
	defer rows.Close()
	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.DateCreated); err != nil {
			logger.Error("error while trying to find by status", err)
			return nil, errors.NewInternalServerError("database error")
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no user matching status %s", status))
	}
	return results, nil

}
