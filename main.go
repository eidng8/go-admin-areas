package main

import (
	"log"
	"os"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/eidng8/go-db"
	"github.com/gin-gonic/gin"

	"github.com/eidng8/go-admin-areas/ent"
	_ "github.com/eidng8/go-admin-areas/ent/runtime"
)

func main() {
	entClient := getEntClient()
	engine, err := NewEngine(
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
		log.Fatalf("Failed to start server: %s", err)
	}
}

func getEntClient() *ent.Client {
	return ent.NewClient(ent.Driver(entsql.OpenDB(db.ConnectX())))
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
