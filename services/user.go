package services

import "auuuth/models"

// AddUser ...
func AddUser(user models.User) error {
	user.AddIndex(user.UserName)

	return user.Store()
}
