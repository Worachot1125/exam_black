package migrations

import "app/app/model"

func Models() []any {
	return []any{
		(*model.User)(nil),
		(*model.Emergency_report)(nil),
		(*model.Emergency_Type)(nil),
		(*model.Role)(nil),
		(*model.User_Role)(nil),
	}
}

func RawBeforeQueryMigrate() []string {
	return []string{
		`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`,
	}
}

func RawAfterQueryMigrate() []string {
	return []string{}
}
