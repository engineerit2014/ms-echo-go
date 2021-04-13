package migrations

import (
	"fmt"
	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
)

func Init(db *pg.DB) {
	// create a new collection with gopg_migrations table
	c := migrations.NewCollection()
	c.DisableSQLAutodiscover(true)
	c.DiscoverSQLMigrations(fmt.Sprintf("migrations"))

	// create gopg_migrations table if not exists
	_, _, _ = c.Run(db, "init")

	// run migration files
	oldVersion, newVersion, err := c.Run(db, "up")
	if err != nil {
		panic(err.Error())
	}
	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
}
