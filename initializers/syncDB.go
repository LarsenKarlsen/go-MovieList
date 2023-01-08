package initializers

import "movieList/models"

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}
