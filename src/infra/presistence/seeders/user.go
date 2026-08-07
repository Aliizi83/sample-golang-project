package seeders

import (
	constant "github.com/Aliizi83/sample-golang-project/src/constants"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userSeed struct {
	userModel    *models.User
	defaultRoles string
}

var defaultUsers = []*userSeed{
	{userModel: &models.User{Username: constant.DefaultUserUsername, Email: constant.DefaultUserEmail, Password: constant.DefaultUserPassword},
		defaultRoles: constant.DefaultAdminRoleTitle},

	{userModel: &models.User{Username: constant.DefaultCustomerUsername, Email: constant.DefaultCustomerEmail, Password: constant.DefaultCustomerPassword},
		defaultRoles: constant.DefaultCustomerRoleTitle},
}

func seedUsers(database *gorm.DB) error {
	count := 0

	database.Model(&models.User{}).Select("COUNT(*)").Find(&count)

	if count == 0 {
		for _, user := range defaultUsers {
			var currentUser models.User
			var defaultRole models.Role

			bytePass, err := bcrypt.GenerateFromPassword([]byte(user.userModel.Password), bcrypt.DefaultCost)
			if err != nil {
				return err
			}

			user.userModel.Password = string(bytePass)

			err = database.Where(models.User{Username: user.userModel.Username}).
				FirstOrCreate(&currentUser, user.userModel).Error
			if err != nil {
				return err
			}

			err = database.Where("title = ?", user.defaultRoles).First(&defaultRole).Error
			if err != nil {
				return err
			}

			var userRole models.UserRole
			err = database.Where("user_id = ? AND role_id = ?", currentUser.Id, defaultRole.Id).First(&userRole).Error
			if err == gorm.ErrRecordNotFound {
				userRole = models.UserRole{
					UserId: currentUser.Id,
					RoleId: defaultRole.Id,
				}
				if err := database.Create(&userRole).Error; err != nil {
					return err
				}
			} else if err != nil {
				return err
			}
		}
	}

	return nil
}
