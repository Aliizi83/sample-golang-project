package seeders

import (
	constant "github.com/Aliizi83/sample-golang-project/src/constants"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"gorm.io/gorm"
)

var defaultRoles = []*models.Role{
	{Title: constant.DefaultAdminRoleTitle, Name: constant.DefaultAdminRoleTitle},
	{Title: constant.DefaultCustomerRoleTitle, Name: constant.DefaultCustomerRoleTitle},
}

func AddDefaultRoles(database *gorm.DB) error {
	count := 0
	database.Model(&models.Role{}).Select("COUNT(*)").Find(&count)

	if count == 0 {
		for _, role := range defaultRoles {
			var existingRole models.Role
			err := database.Where("title = ?", role.Title).First(&existingRole).Error

			if err == gorm.ErrRecordNotFound {
				if err := database.Create(role).Error; err != nil {
					return err
				}
			} else if err != nil {
				return err
			}
		}
	}

	return nil
}
