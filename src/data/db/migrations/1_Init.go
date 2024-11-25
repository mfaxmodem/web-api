package migrations

import (
	"github.com/mfaxmodem/web-api/src/config"
	"github.com/mfaxmodem/web-api/src/constants"
	"github.com/mfaxmodem/web-api/src/data/db"
	"github.com/mfaxmodem/web-api/src/data/models"
	"github.com/mfaxmodem/web-api/src/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Up_1() {
	// Get configuration
	cfg, err := config.GetConfig()
	var logger = logging.NewLogger(cfg)
	if err != nil {
		panic("Failed to get config: " + err.Error())
	}

	// Get database connection
	database := db.GetDb()

	// Create tables
	createTables(database, logger)
	createDefaultInformation(database)
}

func createTables(database *gorm.DB, logger logging.Logger) {
	tables := []interface{}{}
	// User
	tables = addNewTable(database, models.User{}, tables)
	tables = addNewTable(database, models.Role{}, tables)
	tables = addNewTable(database, models.UserRole{}, tables)

	err := database.Migrator().CreateTable(tables...)
	if err != nil {
		logger.Error(logging.Postgres, logging.Migration, err.Error(), nil)
	}
	logger.Info(logging.Postgres, logging.Migration, "tables created", nil)
}

func addNewTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}

func createDefaultInformation(database *gorm.DB) {
	adminRole := models.Role{Name: constants.AdminRoleName}
	createRoleNotExists(database, &adminRole)

	defaultRole := models.Role{Name: constants.DefaultRoleName}
	createRoleNotExists(database, &defaultRole)

	user := models.User{Username: constants.DefaultUserName, FirstName: "Test", LastName: "Test",
		MobileNumber: "09111112222", Email: "admin@admin.com"}
	pass := "12345678"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	createAdminUserIfNotExists(database, &user, adminRole.Id)
}

func createRoleNotExists(database *gorm.DB, r *models.Role) {
	exists := 0
	database.
		Model(&models.Role{}).
		Select("1").
		Where("name = ?", r.Name).
		First(&exists)
	if exists == 0 {
		database.Create(r)
	}

}

func createAdminUserIfNotExists(database *gorm.DB, user *models.User, roleId int) {
	exists := 0
	database.
		Model(&models.User{}).
		Select("1").
		Where("Username = ?", user.Username).
		First(&exists)
	if exists == 0 {
		database.Create(user)
		user := models.UserRole{UserId: user.Id, RoleId: roleId}
		database.Create(&user)
	}

}
