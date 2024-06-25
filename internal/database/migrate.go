package database

import "github.com/Guilherme-Matosoli/votter/internal/database"

func RunMigration() {
	conn, err := database.Connection()
	if err != nil {
		panic(err)
	}

	sql := `CREATE TABLE `
}
