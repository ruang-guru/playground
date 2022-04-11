package utils

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
)

func ParseError(err error) TodoErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), "no rows in result set") {
			return NewNotFoundError("no record matching given id")
		}
		return NewInternalServerError(fmt.Sprintf("error when trying to save todo: %s", err.Error()))
	}
	switch sqlErr.Number {
	case 1062:
		return NewInternalServerError("title already taken")
	}
	return NewInternalServerError(fmt.Sprintf("error when processing request: %s", err.Error()))
}
