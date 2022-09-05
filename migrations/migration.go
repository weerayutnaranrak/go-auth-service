package migrations

import (
	"auth-service/config"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
)

func Migrate() {
	m := gormigrate.New(
		config.GetDB(),
		gormigrate.DefaultOptions,
		[]*gormigrate.Migration{
			m16620166385User(),
			m1662091994User(),
		},
	)
	if err := m.Migrate(); err != nil {
		log.Fatalf("Cound not migrate %v", err)
	}
}
