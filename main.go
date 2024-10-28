package main

import (
	"database/sql"
	"log"
	"os"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"

	"eidng8.cc/microservices/admin-areas/api"
	"eidng8.cc/microservices/admin-areas/ent"
	_ "eidng8.cc/microservices/admin-areas/ent/runtime"
)

func main() {
	entClient := getEntClient()
	engine, err := api.NewEngine(
		getenvd("SERVER_MODE", gin.ReleaseMode), entClient,
	)
	if err != nil {
		log.Fatalf("Failed to create server: %s", err)
	}
	defer func(entClient *ent.Client) {
		err := entClient.Close()
		if err != nil {
			log.Fatalf("Failed to close ent client: %s", err)
		}
	}(entClient)
	if err := engine.Run(getenv("LISTEN_ADDR")); err != nil {
		log.Fatalf("Failed to getECS server: %s", err)
	}
}

func getEntClient() *ent.Client {
	var db *sql.DB
	drv := getenvd("DB_DRIVER", dialect.MySQL)
	switch {
	case dialect.MySQL == drv:
		db = openMysql()
	case dialect.SQLite == drv:
		db = openSqlite()
	case dialect.Postgres == drv:
		db = openPostgres()
	}
	return ent.NewClient(ent.Driver(entsql.OpenDB(drv, db)))
}

func getenv(name string) string {
	val := os.Getenv(name)
	if "" == val {
		log.Fatalf("%v environment variable is not set", name)
	}
	return val
}

func getenvd(name, defval string) string {
	val := os.Getenv(name)
	if "" == val {
		return defval
	}
	return val
}
