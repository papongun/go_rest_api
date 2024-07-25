package repository_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/papongun/go_todo/repository"
	"github.com/papongun/go_todo/test/db_mock"
)

func TestCreateUserShouldSuccess(t *testing.T) {
	sql, db, mock := db_mock.DbMock(t)
	defer sql.Close()
	repo := *repository.GetUserRepository(db)

	addRow := sqlmock.NewRows([]string{"id"}).AddRow("1")
	expectedSQL := "INSERT INTO \"users\" (.+) VALUES (.+)"
	mock.ExpectBegin()
	mock.ExpectQuery(expectedSQL).WillReturnRows(addRow)
	mock.ExpectCommit()

	username := "papon"
	displayName := "papon"
	password := "awddfgr44f3fewdf3"

	repo.Save(username, displayName, password)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestAddUserDuplicateUsername(t *testing.T) {
	sql, db, mock := db_mock.DbMock(t)
	defer sql.Close()
	repo := *repository.GetUserRepository(db)

	expectedSQL := "INSERT INTO \"users\" (.+) VALUES (.+)"
	mock.ExpectBegin()
	mock.ExpectQuery(expectedSQL).WillReturnError(errors.New("duplicate key value violates unique constraint \"idx_users_username\""))
	mock.ExpectRollback()

	username := "duplicate"
	displayName := "duplicate"
	password := "password"

	_, err := repo.Save(username, displayName, password)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "duplicate key value violates unique constraint \"idx_users_username\"")
	assert.Nil(t, mock.ExpectationsWereMet())
}
