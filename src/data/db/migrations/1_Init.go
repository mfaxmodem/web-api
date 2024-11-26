package migration

import (
	"github.com/mfaxmodem/web-api/src/config"
	"github.com/mfaxmodem/web-api/src/data/db"
	"github.com/mfaxmodem/web-api/src/data/models"
	"github.com/mfaxmodem/web-api/src/pkg/logging"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var logger = initializeLogger()

func pointerToString(s string) *string {
	return &s
}
func initializeLogger() logging.Logger {
	cfg, err := config.GetConfig()
	if err != nil {
		panic("failed to load configuration: " + err.Error()) // مدیریت خطای کانفیگ
	}
	return logging.NewLogger(cfg) // فقط مقدار config به NewLogger ارسال می‌شود
}
func Up1() {
	database := db.GetDb()

	createTables(database)
	createDefaultUserInformation(database)

}

func createTables(database *gorm.DB) {
	tables := []interface{}{
		models.User{},
		models.Role{},
		models.UserRole{},
	}

	for _, table := range tables {
		if !database.Migrator().HasTable(table) {
			if err := database.Migrator().CreateTable(table); err != nil {
				logger.Error(logging.Postgres, logging.Migration, err.Error(), nil)
				return // اگر خطا وجود دارد، ادامه ندهید
			}
		}
	}

	logger.Info(logging.Postgres, logging.Migration, "tables created", nil)
}

func createDefaultUserInformation(database *gorm.DB) {

	adminRole := models.Role{Name: "admin"}
	createRoleIfNotExists(database, &adminRole)

	defaultRole := models.Role{Name: "default"}
	createRoleIfNotExists(database, &defaultRole)

	u := models.User{
		Username:     "admin",
		FirstName:    pointerToString("Test"),            // استفاده از اشاره‌گر
		LastName:     pointerToString("Test"),            // استفاده از اشاره‌گر
		MobileNumber: pointerToString("09120000000"),     // استفاده از اشاره‌گر
		Email:        pointerToString("admin@admin.com"), // استفاده از اشاره‌گر
	}
	pass := "12345678"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	u.Password = string(hashedPassword)

	createAdminUserIfNotExists(database, &u, adminRole.Id)

}

func createRoleIfNotExists(database *gorm.DB, r *models.Role) {
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

func createAdminUserIfNotExists(database *gorm.DB, u *models.User, roleId int) {
	exists := 0
	database.
		Model(&models.User{}).
		Select("1").
		Where("username = ?", u.Username).
		First(&exists)
	if exists == 0 {
		database.Create(u)
		ur := models.UserRole{UserId: u.Id, RoleId: roleId}
		database.Create(&ur)
	}
}

func Down1() {
	// nothing
}
