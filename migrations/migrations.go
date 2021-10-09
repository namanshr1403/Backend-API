package migrations

import (
	"duomly.com/go-bank-backend/helpers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
	Name     string
}

type password struct {
	gorm.Model
	Email    string
	Name     string
	password int
	UserID   uint
}

func connectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=user dbname=dbname password=password sslmode=disable")
	helpers.HandleErr(err)
	return db
}

// This is correct way of creating password
// func HashAndSalt(pass []byte) string {
// 	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
// 	helpers.HandleErr(err)

// 	return string(hashed)
// }

func createAccounts() {
	db := connectDB()

	users := [2]User{
		{Username: "Naman", Email: "namanshrivastava94253@gmail.com"},
		{Username: "Naman", Email: "namanshrivastava94253@gmail.com"},
	}

	for i := 0; i < len(users); i++ {
		// Correct one way
		// generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		generatedPassword := helpers.HashOnlyVulnerable([]byte(users[i].Username))
		user := User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)

		password := password{
			Model: gorm.Model{},

			Name:     "Saket Nand",
			password: 25669,
			UserID:   user.ID,
			Email:    "saketnand@gmail.com",
		}
		db.Create(&password)
	}
	defer db.Close()
}

func Migrate() {
	db := connectDB()
	db.AutoMigrate(&User{}, &password{})
	defer db.Close()

	createAccounts()
}
