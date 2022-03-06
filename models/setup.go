// models/setup.go

package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetConnectionString() string {
	dbuser, dbuserpresent := os.LookupEnv("BLOG_DBUSER")
	dbpass, dbpasspresent := os.LookupEnv("BLOG_DBPASS")
	dbhost, dbhostpresent := os.LookupEnv("BLOG_DBHOST")
	dbname, dbnamepresent := os.LookupEnv("BLOG_DB")
	dbport, dbportpresent := os.LookupEnv("BLOG_DBPORT")

	if !dbuserpresent {
		log.Panicf("Database username is not provided")
	}

	if !dbpasspresent {
		log.Panicf("Database password is not provided")
	}

	if !dbhostpresent {
		log.Panicf("Database hostname is not provided")
	}

	if !dbnamepresent {
		log.Panicf("Database name isn't provided")
	}

	if !dbportpresent {
		log.Printf("Port not provided, using 5432")
		dbport = "5432"
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=EST", dbhost, dbport, dbuser, dbpass, dbname)
}

func ConnectDatabase() {
	db, err := gorm.Open(postgres.Open(GetConnectionString()), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database.")
	}

	db.AutoMigrate(&Blog{})

	DB = db
}
