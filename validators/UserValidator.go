package validators

import (
	"fmt"
	"movies-user/domain/users"
	"movies-user/utils/constants"
	"movies-user/utils/rest_errors"
	"strings"
)

func Validate(user *users.User) rest_errors.RestErr{
	fmt.Println("INSIDE VALIDATE     ")
	user.MiddleName = strings.TrimSpace(user.MiddleName)
	user.FirstName = strings.TrimSpace(strings.ToUpper(user.FirstName))
	if user.FirstName == "" {
		return rest_errors.NewBadRequestError("First Name is mandatory field")
	}
	if user.LastName == "" {
		return rest_errors.NewBadRequestError("Last Name is mandatory field")
	}
	user.Email = strings.TrimSpace(strings.ToUpper(user.Email))
	if user.Email == "" {
		return rest_errors.NewBadRequestError("Email is mandatory field")
	}
	user.UserName = strings.ToUpper(strings.ToUpper(user.UserName))
	if user.UserName == "" {
		return rest_errors.NewBadRequestError("User name is mandatory field")
	}
	passwordLength := len(user.Password)
	if passwordLength<constants.PasswordMinLength || passwordLength > constants.PasswordMaxLength {
		return rest_errors.NewBadRequestError(fmt.Sprintf(`Password length must be between %d and %d `,
			constants.PasswordMinLength,constants.PasswordMaxLength))
	}
	if user.Type == "" {
		return rest_errors.NewBadRequestError("Type is a mandatory field")
	}
	return nil
}