package migrationbackupper

import (
	"fmt"

	"github.com/keighl/barkup"
)

func CreateBuckup() {
	postgresDB := &barkup.Postgres{
		Host:     "postgres",
		Port:     "5432",
		DB:       "postgres",
		Username: "postgres",
		Options:  []string{"--no-owner"},
	}
	backup := postgresDB.Export()
	if backup.Error != nil {
		fmt.Println(backup.Error)
		panic("Could not create a backup!")
	}
	fmt.Println("backup path:", backup.Path)
}

// todo: do backups
// todo: save backup to hard drive
// todo: upload backup to s3 first
// todo: upload backup to sfpt first
