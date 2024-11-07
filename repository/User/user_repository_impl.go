package user

import (
	authentication "Expire/data/request/Authentication"
	"Expire/helper"
	"Expire/model"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

func (t *UserRepositoryImpl) GetUserByID(id string) (model.User, *helper.CustomError) {
	var user model.User

	userId, err := uuid.Parse(id)

	data := t.Db.First(&user, "id = ?", userId)

	if data == nil || err != nil {
		return user, &helper.CustomError{
			Code:    404,
			Message: "User Not Founded.",
		}
	}

	user.Password = ""

	return user, nil
}

func (t *UserRepositoryImpl) Create(user model.User) *helper.CustomError {
	result := t.Db.Create(&user)

	if result.Error != nil {
		fileName, atLine := helper.GetFileAndLine(result.Error)
		return &helper.CustomError{
			Code:     500,
			Message:  "Unexpected Error When Creating New User.",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	return nil
}

func (t *UserRepositoryImpl) ValidateUserAccount(payload authentication.SignInInput) (*model.User, *helper.CustomError) {
	var err error

	user := model.User{}

	if err = t.Db.Model(model.User{}).First(&user, "email = ?", strings.ToLower(payload.Email)).Error; err != nil {
		fileName, atLine := helper.GetFileAndLine(err)
		return nil, &helper.CustomError{
			Code:     403,
			Message:  "The Email hasn't been Registered",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	fmt.Println("User Password", "Payload Password", user.Password, payload.Password)

	err = VerifyPassword(user.Password, payload.Password)

	if err != nil {
		fileName, atLine := helper.GetFileAndLine(err)
		return nil, &helper.CustomError{
			Code:     403,
			Message:  "Username or Password is Wrong",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	return &user, nil
}

func (t *UserRepositoryImpl) FindUserById(userId string) (*model.User, *helper.CustomError) {
	var err error
	user := model.User{}

	if err = t.Db.Model(model.User{}).First(&user, "id = ?", strings.ToLower(userId)).Error; err != nil {
		return nil, &helper.CustomError{
			Code:    404,
			Message: "User Not Founded.",
		}
	}
	return &user, nil
}

func VerifyPassword(hashPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}

func (t *UserRepositoryImpl) FindUserByEmail(email string) (*model.User, *helper.CustomError) {
	var err error

	user := model.User{}

	if err = t.Db.Model(model.User{}).First(&user, "email = ?", strings.ToLower(email)).Error; err != nil {
		return nil, &helper.CustomError{
			Code:    404,
			Message: "User Not Founded.",
		}
	}

	return &user, nil
}

func (t *UserRepositoryImpl) UpdatePasssword(email string, newPassword string) *helper.CustomError {
	var user model.User
	var errFindUser error
	var errUpdateUser error

	if errFindUser = t.Db.Model(model.User{}).First(&user, "email = ?", strings.ToLower(email)).Error; errFindUser != nil {
		return &helper.CustomError{
			Code:    404,
			Message: "User Not Founded.",
		}
	}

	fmt.Println("USER:", user)

	decryptedPassword, errorDecrypt := model.ManualDecryptPassword(newPassword)

	if errorDecrypt != nil {
		fileName, atLine := helper.GetFileAndLine(errorDecrypt)
		return &helper.CustomError{
			Code:     400,
			Message:  "Password Decryption Error.",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	if errUpdateUser = t.Db.Model(&user).UpdateColumn("password", decryptedPassword).Error; errUpdateUser != nil {
		fileName, atLine := helper.GetFileAndLine(errorDecrypt)
		return &helper.CustomError{
			Code:     500,
			Message:  "Unexpected Error When Creating New User.",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	return nil
}

func (t *UserRepositoryImpl) GetAllUser() ([]model.User, *helper.CustomError) {
	var users []model.User
	result := t.Db.Order("created_at DESC").Find(&users)
	if result.Error != nil {
		fileName, atLine := helper.GetFileAndLine(result.Error)
		return nil, &helper.CustomError{
			Code:     500,
			Message:  "Unexpected Error When Fetching Reports",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	return users, nil
}

func (t *UserRepositoryImpl) Delete(id string) *helper.CustomError {
	userId, err := uuid.Parse(id)
	result := t.Db.Unscoped().Delete(&model.User{}, userId)

	if result.Error != nil || err != nil {
		fileName, atLine := helper.GetFileAndLine(result.Error)
		return &helper.CustomError{
			Code:     500,
			Message:  "Unexpected Error When Fetching Reports",
			FileName: fileName,
			AtLine:   atLine,
		}
	}

	return nil
}
