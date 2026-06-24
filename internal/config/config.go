package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DatabaseURL        string
	JWTSecret          string
	RefreshTokenSecret string
	AppPort            string
}

var Environ Env

// Load يقرأ المتغيرات من البيئة ويتحقق من وجود المطلوب منها.
// يرجّع error بدل ما يوقف البرنامج بنفسه، عشان main.go هي اللي تتحكم.
func Load() (Env, error) {
	// لو الملف مش موجود ده مش خطأ — في الإنتاج المتغيرات بتيجي من النظام.
	if err := godotenv.Load(); 
	err != nil {
		fmt.Fprintln(os.Stderr, "Note: no .env file found, reading from system environment")
	}

	required := []string{"DATABASE_URL", "JWT_SECRET", "REFRESH_TOKEN_SECRET"}
	for _, key := range required {
		if os.Getenv(key) == "" {
			return Env{}, fmt.Errorf("required environment variable %s is not set", key)
		}
	}

	env := Env{
		DatabaseURL:        os.Getenv("DATABASE_URL"),
		JWTSecret:          os.Getenv("JWT_SECRET"),
		RefreshTokenSecret: os.Getenv("REFRESH_TOKEN_SECRET"),
		AppPort:            getEnvOr("APP_PORT", "8080"), // اختياري بقيمة افتراضية
	}

	Environ = env
	return env, nil
}

// getEnvOr يرجّع قيمة المتغير، أو القيمة الافتراضية لو فاضي.
func getEnvOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}