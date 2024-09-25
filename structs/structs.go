package structs

import (
	"fmt"

	"github.com/adnux/go-basic-projects/structs/user"
)

func StartUserStructs() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthday := getUserData("Please enter your birthday (DD/MM/YYYY): ")

	var appUser *user.User

	appUser, err := user.New(userFirstName, userLastName, userBirthday)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

	adminUser, errAdmin := user.NewAdmin(
		"email",
		"password",
		userFirstName,
		userLastName,
		userBirthday,
	)
	if errAdmin != nil {
		fmt.Println(errAdmin)
		return
	}
	adminUser.OutputAdminDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
