
package migrations

import (
	"github.com/guswitch/spiderman-api.git/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Character{})
}