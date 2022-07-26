package database

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"gebes.io/sticker_backend/pkg/ent"
	"gebes.io/sticker_backend/pkg/env"
	"gebes.io/sticker_backend/pkg/logger"
	_ "github.com/lib/pq"

	"time"
)

var (
	Client *ent.Client
	Driver *sql.Driver
)

func init() {
	logger.Info.Println("Initializing Postgres database connection")
	var err error
	Driver, err = sql.Open("postgres", env.PostgresDatabase)
	if err != nil {
		logger.Error.Fatalln("Failed opening connection to Postgres:", err)
	}
	// Get the underlying sql.DB object of the driver.
	db := Driver.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	Client = ent.NewClient(ent.Driver(Driver)) //.Debug()
	err = autoMigrate()
	if err != nil {
		logger.Error.Fatalln("Failed to auto migrate database:", err)
	}
	logger.Info.Println("Migrated Postgres database")
}

func autoMigrate() error {
	return Client.Schema.Create(context.Background())
}

func rollback(errToReturn error, tx *ent.Tx) error {
	err := tx.Rollback()
	if err != nil {
		return err
	}
	return errToReturn
}

func Ping() (time.Duration, error) {
	start := time.Now()
	err := Driver.DB().Ping()
	return time.Now().Sub(start), err
}

func Close() error {
	return Client.Close()
}
