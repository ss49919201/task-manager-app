package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_run(t *testing.T) {
	// FIXME: テスト用Dockerコンテナを起動する
	t.SkipNow()

	t.Setenv("DATABASE_HOST", "127.0.0.1")
	t.Setenv("DATABASE_PORT", "3306")
	t.Setenv("DATABASE_USER", "user")
	t.Setenv("DATABASE_PASS", "password")
	t.Setenv("DATABASE_NAME", "database")

	require.NoError(t, run())
}
