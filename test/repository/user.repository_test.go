package repository_test

import (
	"bookstore-framework/internal/users"
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	require.NoError(t, err)

	return gormDB, mock
}

func TestRegisterRepository_Success(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	gormDB, mock := setupMockDB(t)
	repo := users.NewUserRepository(gormDB)

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "users"`).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	testUser := &users.User{
		Username: "testuser",
		Email:    "test@gmail.com",
		Password: "password123",
	}

	result, err := repo.Register(context.Background(), testUser)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "testuser", result.Username)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)

}
