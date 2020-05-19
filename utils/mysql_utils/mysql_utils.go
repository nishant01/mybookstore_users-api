package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/nishant01/mybookstore_users-api/utils/errors"
)

const (
	// INDEX_UNIQUE_EMAIL = "email_UNIQUE"
	NO_ROW_IN_RESULT = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), NO_ROW_IN_RESULT) {
			return errors.NewNotFoundError("No record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("Email already exist.")
	}
	return errors.NewInternalServerError("error processing request")
}
