package main

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	jitr "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"

	"eidng8.cc/microservices/admin-areas/ent"
)

var jsoniter = jitr.ConfigCompatibleWithStandardLibrary

func setupGinTest(tb testing.TB) (
	*gin.Engine, *ent.Client, *httptest.ResponseRecorder,
) {
	assert.Nil(tb, os.Setenv("DB_DRIVER", "mysql"))
	assert.Nil(tb, os.Setenv("DB_USER", "root"))
	assert.Nil(tb, os.Setenv("DB_PASSWORD", "123456"))
	assert.Nil(tb, os.Setenv("DB_HOST", "127.0.0.1:43306"))
	assert.Nil(tb, os.Setenv("DB_NAME", "admin_areas"))
	entClient := getEntClient()
	tb.Cleanup(
		func() {
			assert.Nil(tb, entClient.Close())
		},
	)
	engine, err := NewEngine(gin.TestMode, entClient)
	assert.Nil(tb, err)
	return engine, entClient, httptest.NewRecorder()
}
