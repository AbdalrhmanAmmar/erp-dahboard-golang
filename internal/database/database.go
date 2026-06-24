package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB المتغير العام اللي كل الموديولات هتستخدمه للوصول للداتابيز.
var DB *gorm.DB

// Connect يفتح الاتصال بالـ Postgres ويظبط الـ connection pool.
func Connect(databaseURL string) error {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // يطبع الـ SQL في التطوير
	})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// الوصول لطبقة sql.DB الأساسية عشان نظبط الـ pool.
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	// إعدادات الـ connection pool — تتحكم في عدد الاتصالات وعمرها.
	sqlDB.SetMaxOpenConns(25)                 // أقصى عدد اتصالات مفتوحة
	sqlDB.SetMaxIdleConns(5)                  // اتصالات جاهزة مستنية
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // عمر الاتصال قبل ما يتجدد

	DB = db
	log.Println("✅ Database connected successfully")
	return nil
}