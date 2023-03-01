package migrations

import (
	repository "auth-service/repository/user"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m1662366357User() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1662366357",
		Migrate: func(d *gorm.DB) error {
			return d.AutoMigrate(&repository.User{})
		},
		Rollback: func(d *gorm.DB) error {
			return d.Migrator().DropTable("users")
		},
	}
}
