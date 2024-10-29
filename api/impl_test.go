package api

import (
	"database/sql"
	"log"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	jitr "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"

	"eidng8.cc/microservices/admin-areas/ent"
	_ "eidng8.cc/microservices/admin-areas/ent/runtime"
)

var jsoniter = jitr.ConfigCompatibleWithStandardLibrary

func setupGinTest(tb testing.TB) (
	*gin.Engine, *ent.Client, *httptest.ResponseRecorder,
) {
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_ADDR"),
		DBName:               os.Getenv("DB_NAME"),
		MultiStatements:      true,
		AllowNativePasswords: true,
		ParseTime:            true,
		Collation:            "utf8mb4_unicode_ci",
		Loc:                  time.Local,
	}
	db, err := sql.Open(dialect.MySQL, cfg.FormatDSN())
	if err != nil {
		log.Fatalf("Failed to open MySQL connection: %s", err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	entClient := ent.NewClient(ent.Driver(entsql.OpenDB("mysql", db)))
	tb.Cleanup(
		func() {
			if err := entClient.Close(); err != nil {
				tb.Fatalf("Failed to close ent client: %s", err)
			}
		},
	)

	engine, err := NewEngine(gin.TestMode, entClient)
	if err != nil {
		assert.Nil(tb, err)
	}

	res := httptest.NewRecorder()

	return engine, entClient, res
}
