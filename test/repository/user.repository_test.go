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

func TestUserRepository_Success(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	gormDB, mock := setupMockDB(t)
	repo := users.NewUserRepository(gormDB)
	t.Run("Register", func(t *testing.T) {
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
	})

	t.Run("FindUserByUsername", func(t *testing.T) {
		username := "testuser"
		columns := []string{"id", "username", "name", "email", "password"}
		mock.ExpectBegin()
		mock.ExpectQuery(`"SELECT .+ FORM "users" WHERE" username = .+`).
			WithArgs(username).
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(1, "test", "testuser", "test@example.com", "hashedpassword"))

		user, err := repo.FindUserByUsername(context.Background(), username)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, username, user.Username)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)

	})

}
