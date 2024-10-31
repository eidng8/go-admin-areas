package main

import (
	"context"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	jitr "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"

	"github.com/eidng8/go-admin-areas/ent"
)

var jsoniter = jitr.ConfigCompatibleWithStandardLibrary

func setupGinTest(tb testing.TB) (
	*gin.Engine, *ent.Tx, *httptest.ResponseRecorder,
) {
	assert.Nil(tb, os.Setenv("DB_DRIVER", "mysql"))
	assert.Nil(tb, os.Setenv("DB_USER", "root"))
	assert.Nil(tb, os.Setenv("DB_PASSWORD", "123456"))
	assert.Nil(tb, os.Setenv("DB_HOST", "127.0.0.1:43306"))
	assert.Nil(tb, os.Setenv("DB_NAME", "admin_areas"))
	entClient := getEntClient()
	tx, err := entClient.BeginTx(context.Background(), nil)
	assert.Nil(tb, err)
	tb.Cleanup(
		func() {
			_ = tx.Rollback()
			_ = entClient.Close()
		},
	)
	engine, err := NewEngine(gin.TestMode, tx.Client())
	assert.Nil(tb, err)
	return engine, tx, httptest.NewRecorder()
}
