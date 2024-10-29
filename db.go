package main

import (
	"database/sql"
	"log"
	"time"

	"entgo.io/ent/dialect"
	"github.com/go-sql-driver/mysql"
)

func openMysql() *sql.DB {
	cfg := mysql.Config{
		User:                 getenv("DB_USER"),
		Passwd:               getenv("DB_PASSWORD"),
		Net:                  getenvd("DB_PROTOCOL", "tcp"),
		Addr:                 getenv("DB_ADDR"),
		DBName:               getenv("DB_NAME"),
		MultiStatements:      true,
		AllowNativePasswords: true,
		ParseTime:            true,
		Collation:            getenvd("DB_COLLATION", "utf8mb4_unicode_ci"),
		Loc:                  time.Local,
	}
	db, err := sql.Open(dialect.MySQL, cfg.FormatDSN())
	if err != nil {
		log.Fatalf("Failed to open MySQL connection: %s", err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	return db
}

func openSqlite() *sql.DB {
	dsn := getenv("DB_PATH")
	db, err := sql.Open(dialect.SQLite, dsn)
	if err != nil {
		log.Fatalf("Failed to open SQLite connection: %s", err)
	}
	return db
}

func openPostgres() *sql.DB {
	url := getenv("DB_URL")
	db, err := sql.Open("pgx", url)
	if err != nil {
		log.Fatalf("Failed to open PostgeSQL connection: %s", err)
	}
	return db
}
