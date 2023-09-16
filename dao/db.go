package dao

import (
	"context"
)

func GetAllTables(ctx context.Context) ([]string, error) {
	return db.Migrator().GetTables()
}

func DDLExe(ctx context.Context, sql string) error {
	return db.Exec(sql).Error
}
