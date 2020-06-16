package services

import (
	"auuuth/models"
)

// AddUser ...
func AddUser(user models.User) error {
	record := user.IntoRecord()

	return record.Commit()
}
