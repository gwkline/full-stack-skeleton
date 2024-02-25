package repo

import (
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gwkline/full-stack-skeleton/backend/types"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type repoTest struct {
	name          string
	input         *types.User
	mockSetup     func(mock sqlmock.Sqlmock)
	expectedError string
}

func newMockDB() (*Repository, sqlmock.Sqlmock) {
	sql, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("Failed to create sqlmock: %s", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sql,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to create sqlmock: %s", err)
	}

	mockDatabase := NewRepository(gormDB)
	return mockDatabase, mock
}

func TestCreate(t *testing.T) {
	tests := []repoTest{
		{
			name:  "successful creation",
			input: &types.User{Email: "john@example.com"},
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "email"}).
					AddRow(1, "john@example.com")

				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
					WillReturnRows(rows)
				mock.ExpectCommit()
			},
			expectedError: "",
		},
		{
			name:  "error on creation",
			input: &types.User{Email: "jane@example.com"},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
					WillReturnError(gorm.ErrInvalidTransaction)
				mock.ExpectRollback()
			},
			expectedError: "invalid transaction",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()

			db, mock := newMockDB()
			tt.mockSetup(mock)

			result, err := db.User.Create(tt.input)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := []repoTest{
		{
			name:  "successful update",
			input: &types.User{BaseModel: types.BaseModel{ID: 1}, Email: "john_updated@example.com"},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users"`)).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectedError: "",
		},
		{
			name:  "error on update",
			input: &types.User{BaseModel: types.BaseModel{ID: 2}, Email: "jane_failure@example.com"},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users"`)).
					WillReturnError(gorm.ErrInvalidTransaction)
				mock.ExpectRollback()
			},
			expectedError: "invalid transaction",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()

			db, mock := newMockDB()
			tt.mockSetup(mock)

			result, err := db.User.Update(tt.input)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestArchive(t *testing.T) {
	tests := []repoTest{
		{
			name:  "successful archive",
			input: &types.User{BaseModel: types.BaseModel{ID: 1}},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users"`)).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectedError: "",
		},
		{
			name:  "error on archive",
			input: &types.User{BaseModel: types.BaseModel{ID: 2}},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users"`)).
					WillReturnError(gorm.ErrInvalidTransaction)
				mock.ExpectRollback()
			},
			expectedError: "invalid transaction",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()

			db, mock := newMockDB()
			tt.mockSetup(mock)

			err := db.User.Archive(tt.input)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []repoTest{
		{
			name:  "successful delete",
			input: &types.User{BaseModel: types.BaseModel{ID: 1}},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "users"`)).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			expectedError: "",
		},
		{
			name:  "error on delete",
			input: &types.User{BaseModel: types.BaseModel{ID: 2}},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "users"`)).
					WillReturnError(gorm.ErrInvalidTransaction)
				mock.ExpectRollback()
			},
			expectedError: "invalid transaction",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()

			db, mock := newMockDB()
			tt.mockSetup(mock)

			err := db.User.Delete(tt.input)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestListBy(t *testing.T) {
	tests := []struct {
		name          string
		filters       []types.Filter
		mockSetup     func(mock sqlmock.Sqlmock)
		expectedError string
	}{
		{
			name: "successful list by",
			filters: []types.Filter{
				{Key: "email", Value: "john@example.com"},
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "email"}).
					AddRow(1, "john@example.com")
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1 AND "users"."deleted_at" IS NULL`)).
					WillReturnRows(rows)
			},
			expectedError: "",
		},
		{
			name: "error on list by",
			filters: []types.Filter{
				{Key: "email", Value: "jane@example.com"},
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1 AND "users"."deleted_at" IS NULL`)).
					WillReturnError(gorm.ErrInvalidTransaction)
			},
			expectedError: "invalid transaction",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()

			db, mock := newMockDB()
			tt.mockSetup(mock)

			result, err := db.User.ListBy(tt.filters)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, result)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestFindBy(t *testing.T) {
	tests := []struct {
		name          string
		filters       []types.Filter
		mockSetup     func(mock sqlmock.Sqlmock)
		expectedError string
	}{
		{
			name: "successful find by",
			filters: []types.Filter{
				{Key: "id", Value: 1},
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "email"}).
					AddRow(1, "john@example.com")
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE id = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT $2`)).
					WillReturnRows(rows)
			},
			expectedError: "",
		},
		{
			name: "error on find by",
			filters: []types.Filter{
				{Key: "id", Value: 2},
			},
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE id = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT $2`)).
					WillReturnError(gorm.ErrInvalidTransaction)
			},
			expectedError: "invalid transaction",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()

			db, mock := newMockDB()
			tt.mockSetup(mock)

			result, err := db.User.FindBy(tt.filters)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestFuzzyFindBy(t *testing.T) {
	tests := []struct {
		name          string
		key           string
		value         any
		mockSetup     func(mock sqlmock.Sqlmock)
		expectedError string
	}{
		{
			name:  "successful fuzzy find by",
			key:   "email",
			value: "example",
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "email"}).
					AddRow(1, "john@example.com").
					AddRow(2, "jane@example.com")
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE email LIKE $1 AND "users"."deleted_at" IS NULL`)).
					WithArgs("%example%").
					WillReturnRows(rows)
			},
			expectedError: "",
		},
		{
			name:  "error on fuzzy find by",
			key:   "email",
			value: "doesnotexist",
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE email LIKE $1 AND "users"."deleted_at" IS NULL`)).
					WithArgs("%doesnotexist%").
					WillReturnError(gorm.ErrInvalidTransaction)
			},
			expectedError: "invalid transaction",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()

			db, mock := newMockDB()
			tt.mockSetup(mock)

			result, err := db.User.FuzzyFindBy(tt.key, tt.value)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, result)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
