package repository

import (
	"context"
	"database/sql"
	"fmt"
	"ginweb/domain"
	"ginweb/ent"
	"testing"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/stretchr/testify/assert"
)

// GetUsersのテストコード
func TestGetUsers(t *testing.T) {
	// テスト用のユーザーデータを作成
	mockUsers := []domain.User{
		{
			ID:   1,
			Name: "aaa",
			Age:  20,
		},
		{
			ID:   2,
			Name: "bbb",
			Age:  21,
		},
		{
			ID:   3,
			Name: "ccc",
			Age:  22,
		},
	}

	client, err := Open()
	if err != nil {
		t.Fatalf("failed to get users: %v", err)
	}

	// userRepositoryのモックを作成
	mockRepo := &userRepository{
		DB: client, // モックのDBインスタンスを設定する
	}

	// モックのDBインスタンスを使ってテストを実行
	users, err := mockRepo.GetUsers(context.Background())

	// テスト結果の検証
	assert.NoError(t, err)            // エラーがないことを確認
	assert.Equal(t, mockUsers, users) // ユーザーが期待通りであることを確認
}

func TestGetUserByID(t *testing.T) {
	// テスト用のユーザーデータを作成
	mockUser := domain.User{
		ID:   4,
		Name: "kkk",
		Age:  24,
	}

	client, err := Open()
	if err != nil {
		t.Fatalf("failed to get users: %v", err)
	}

	// userRepositoryのモックを作成
	mockRepo := &userRepository{
		DB: client, // モックのDBインスタンスを設定する
	}

	// モックのDBインスタンスを使ってテストを実行
	user, err := mockRepo.GetUserByID(context.Background(), 4)

	// テスト結果の検証
	assert.NoError(t, err)          // エラーがないことを確認
	assert.Equal(t, mockUser, user) // ユーザーが期待通りであることを確認
}

func Open() (*ent.Client, error) {
	user := "postgres"
	password := "sa"
	host := "localhost"
	port := "5555"
	dbname := "template"
	dsn := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		err = fmt.Errorf("[infrastructure.Open] failed to open connection to database: %w", err)
		return nil, err
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), nil
}
