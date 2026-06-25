package database

import "github.com/AbdalrhmanAmmar/erp-dahboard-golang/internal/modules/users"

func Migrate() error {
return DB.AutoMigrate(
	&users.User{},
	&users.Role{},
)
}