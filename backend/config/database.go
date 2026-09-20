package config

import (
	model "backend/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	var err error
	// dsn := os.Getenv("DB_DSN")

	// dsn := "postgresql://cv_soundsolid:a2bbccc042f9f1d224ca35e47bb73f5e56db4019@ypjpzm.h.filess.io:5434/cv_soundsolid?search_path=public"
	dsn := "postgresql://neondb_owner:npg_A7x6DHdEezpT@ep-flat-sky-b5fp49ue-pooler.c-7.us-east-2.aws.neon.tech/cv?sslmode=require&channel_binding=require"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database. Error: ", err)
	}

	err = db.AutoMigrate(
		&model.Category{},
	)

	if err != nil {
		log.Fatal("Error Migrating model. Error: ", err)
	}

	return db
}
