package Login

import (
	"errors"

	"blog/database"
	"blog/model"
)

func Authenticate(email, password string) (*model.User, error) {

	db := database.Get_GormDB()

	var user model.User

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("User not found")
	}

	//comparing the provided passowrd with the hashed password in the database
	// err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	// if err != nil {
	// 	return nil, errors.New("Incorrect passowrd")
	// }
	if user.Password != password {
		return nil, errors.New("incorrect password , cant log in")
	}

	return &user, nil
}
