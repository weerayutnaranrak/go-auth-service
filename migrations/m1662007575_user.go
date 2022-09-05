package migrations

import (
	"auth-service/repository"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func m16620166385User() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1662016638",
		Migrate: func(d *gorm.DB) error {
			return d.AutoMigrate(&repository.User{})
		},
		Rollback: func(d *gorm.DB) error {
			return d.Migrator().DropTable("users")
		},
	}
}
